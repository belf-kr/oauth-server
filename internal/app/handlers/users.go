package handlers

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/belf-kr/oauth-server/internal/app/data/orm"
	"github.com/belf-kr/oauth-server/internal/app/entitys"
	"github.com/belf-kr/oauth-server/internal/app/models"
	"github.com/belf-kr/oauth-server/internal/pkg/auth"
	"github.com/belf-kr/oauth-server/internal/pkg/kakao"
	"github.com/belf-kr/oauth-server/internal/pkg/util"
	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	"github.com/golang-jwt/jwt"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

// @Summary 사용자 생성 (회원가입)
// @Description 사용자를 생성합니다.
// @Tags User
// @Accept json
// @Produce json
// @Param data body models.UserSignUp true "회원가입을 위한 정보"
// @Success 201 {object} entitys.User
// @Header 201 {string} Location "/users/1"
// @Failure 400 {object} models.ErrResponse
// @Failure 409 {object} models.ErrResponse
// @Failure 500 {object} models.ErrResponse
// @Router /users [post]
func UserSignup(c *gin.Context) {
	userBody := models.UserSignUp{}

	if err := c.BindJSON(&userBody); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrResponse{
			Message: err.Error(),
		})
		return
	}

	if userBody.Password != userBody.PasswordConfirm {
		c.JSON(http.StatusBadRequest, models.ErrResponse{
			Message: "비밀번호가 일치하지 않습니다.",
		})
		return
	}

	user := entitys.User{
		Email:    userBody.Email,
		Password: userBody.Password,
		Name:     userBody.Name,
	}
	if err := orm.Client.Select("Email", "Password", "Name", "ConnectedAt").Create(&user).Error; err != nil {
		// 중복 키 에러 처리: https://github.com/go-gorm/gorm/issues/4037#issuecomment-771499867
		// 예시 에러 => Error 1062: Duplicate entry 'user01@test.com' for key 'email'
		var mysqlErr *mysql.MySQLError
		if errors.As(err, &mysqlErr) && mysqlErr.Number == 1062 {
			parseKey := strings.Split(mysqlErr.Message, "key ")[1]
			duplicateKey := strings.Replace(parseKey, "'", "", -1)
			switch duplicateKey {
			case "email":
				c.JSON(http.StatusConflict, models.ErrResponse{
					Message: "이미 사용중인 이메일입니다.",
				})
				return
			}
		}

		c.JSON(http.StatusInternalServerError, models.ErrResponse{
			Message: err.Error(),
		})
		return
	}

	c.Header("Location", fmt.Sprintf("/users/%d", user.Id))
	c.JSON(http.StatusCreated, user)
}

// @Summary 로그인
// @Description 로그인을 성공 시 JWT token이 발급됩니다.
// @Tags User
// @Accept json
// @Produce json
// @Param data body models.UserLogin true "로그인 정보"
// @Success 201 {object} models.JWTToken
// @Failure 400 {object} models.ErrResponse
// @Failure 401 {object} models.ErrResponse
// @Failure 404
// @Failure 422 {object} models.ErrResponse
// @Failure 500 {object} models.ErrResponse
// @Router /users/login [post]
func UserLogin(c *gin.Context) {
	userBody := models.UserLogin{}

	if err := c.BindJSON(&userBody); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrResponse{
			Message: err.Error(),
		})
		return
	}

	// 존재하는 사용자 인지 확인
	user := entitys.User{}
	if err := orm.Client.Where("email = ?", userBody.Email).Find(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrResponse{
			Message: err.Error(),
		})
		return
	}
	if userBody.Email != user.Email {
		c.Status(http.StatusNotFound)
		return
	}

	// 비밀번호 일치하는지 확인
	if util.Sha256(userBody.Password) != user.Password {
		c.JSON(http.StatusUnauthorized, models.ErrResponse{
			Message: "비밀번호가 일치하지 않습니다.",
		})
		return
	}

	// 토큰 생성
	userIdUint64 := uint64(user.Id)
	ts, err := auth.CreateSession(userIdUint64)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, models.ErrResponse{
			Message: err.Error(),
		})
		return
	}
	tokens := models.JWTToken{
		AccessToken:  ts.AccessToken,
		RefreshToken: ts.RefreshToken,
	}
	c.JSON(http.StatusCreated, tokens)
}

// @Summary 카카오 로그인
// @Description 카카오 oauth 인증 시 콜백되는 API 입니다. 환경변수에 주입된 경로로 token을 queryString에 실어서 리디렉션 합니다.
// @Tags User
// @Accept json
// @Produce json
// @Param code query string false "카카오 로그인 성공 시 전달되는 인가 코드"
// @Success 302
// @Header 302 {string} Location "http://localhost:3000/auth-end?accessToken=header.payLoad.signature&refreshToken=header.payLoad.signature"
// @Failure 422 {object} models.ErrResponse
// @Failure 500 {object} models.ErrResponse
// @Router /users/login/kakao [get]
func UserKakaoLoginCallBack(c *gin.Context) {
	code := c.Query("code")
	if code == "" {
		c.JSON(http.StatusInternalServerError, models.ErrResponse{
			Message: "인가 코드가 전달되지 않았습니다. url에 있는 error_description key의 value를 확인해 주세요.",
		})
		return
	}

	kakaoToken := models.KakaoToken{}
	if err := kakao.GetToken(code, &kakaoToken); err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrResponse{
			Message: err.Error(),
		})
		return
	}

	kakaoUserInformation := models.KakaoUserInformation{}
	if err := kakao.GetUserInfo(kakaoToken.AccessToken, &kakaoUserInformation); err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrResponse{
			Message: err.Error(),
		})
		return
	}

	// 사용자 테이블에 로그인된 카카오 아이디로 생성된 사용지 있는지 확인
	// TODO: 나중에 조인 쿼리 날려보기
	user := entitys.User{}
	if err := orm.Client.Where("kakao_talk_socials_id = ?", kakaoUserInformation.Id).Find(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrResponse{
			Message: err.Error(),
		})
		return
	}

	if kakaoUserInformation.Id != user.KakaoTalkSocialsId {
		// 카카오 oauth가 처음인 경우: 카카오 소셜 테이블에 생성 후 사용자 테이블에 인설트 후 해당 아이디 응답
		kakaoTalkSocial := entitys.KakaoTalkSocial{
			Id:                kakaoUserInformation.Id,
			Email:             kakaoUserInformation.KakaoAccount.Email,
			NickName:          kakaoUserInformation.KakaoAccount.Profile.Nickname,
			ProfileImageUrl:   kakaoUserInformation.KakaoAccount.Profile.ProfileImageUrl,
			ThumbnailImageUrl: kakaoUserInformation.KakaoAccount.Profile.ThumbnailImageUrl,
		}
		if err := orm.Client.Create(&kakaoTalkSocial).Error; err != nil {
			c.JSON(http.StatusInternalServerError, models.ErrResponse{
				Message: err.Error(),
			})
			return
		}

		// kakao 프로필 사진 바이너리로 전환
		url := kakaoTalkSocial.ProfileImageUrl
		res, err := http.Get(url)
		if err != nil {
			c.JSON(http.StatusInternalServerError, models.ErrResponse{
				Message: err.Error(),
			})
			return
		}
		avatarImageBinary, err := ioutil.ReadAll(res.Body)
		if err != nil {
			c.JSON(http.StatusInternalServerError, models.ErrResponse{
				Message: err.Error(),
			})
			return
		}
		res.Body.Close()

		user = entitys.User{
			Email:              kakaoTalkSocial.Email,
			Name:               kakaoTalkSocial.NickName,
			AvatarImage:        avatarImageBinary,
			KakaoTalkSocialsId: kakaoTalkSocial.Id,
		}
		if err := orm.Client.Select("Email", "Name", "AvatarImage", "KakaoTalkSocialsId").Create(&user).Error; err != nil {
			c.JSON(http.StatusInternalServerError, models.ErrResponse{
				Message: err.Error(),
			})
			return
		}
	}

	// 토큰 생성
	userIdUint64 := uint64(user.Id)
	ts, err := auth.CreateSession(userIdUint64)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, models.ErrResponse{
			Message: err.Error(),
		})
		return
	}
	jwtToken := models.JWTToken{
		AccessToken:  ts.AccessToken,
		RefreshToken: ts.RefreshToken,
	}
	redUrl := auth.AuthSuccessRedUrl(jwtToken)
	c.Redirect(http.StatusFound, redUrl)
}

// @Summary 사용자 조회
// @Description token 클레임에 있는 id 값으로 사용자를 조회합니다.
// @Tags User
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer {AccessToken}"
// @Success 200 {object} models.UserInfo
// @Failure 401 {object} models.ErrResponse
// @Failure 404
// @Failure 500 {object} models.ErrResponse
// @Router /users [get]
func UserInfoTokenQuey(c *gin.Context) {
	au, err := auth.ExtractTokenMetadata(c.Request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, models.ErrResponse{
			Message: err.Error(),
		})
		return
	}

	user := entitys.User{
		Id: uint(au.UserId),
	}
	if err := orm.Client.First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.Status(http.StatusNotFound)
			return
		}

		c.JSON(http.StatusInternalServerError, models.ErrResponse{
			Message: err.Error(),
		})
		return
	}

	res := models.NewUserInfo(user)
	c.JSON(http.StatusOK, res)
}

// @Summary 사용자 조회
// @Description id 혹은 email로 사용자를 조회합니다.
// @Tags User
// @Accept json
// @Produce json
// @Param userKey path string true "id or email"
// @Success 200 {object} models.UserInfo
// @Failure 400
// @Failure 404
// @Failure 500 {object} models.ErrResponse
// @Router /users/{userKey} [get]
func UserInfoQuey(c *gin.Context) {
	userKey := c.Param("userKey")
	if userKey == "" {
		c.Status(http.StatusBadRequest)
		return
	}

	user := entitys.User{}
	result := orm.Client.Where("id = ?", userKey).Or("email = ?", userKey).Find(&user)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, models.ErrResponse{
			Message: result.Error.Error(),
		})
		return
	}
	if result.RowsAffected == 0 {
		c.Status(http.StatusNotFound)
		return
	}

	res := models.NewUserInfo(user)
	c.JSON(http.StatusOK, res)
}

// @Summary 로그아웃
// @Description token 클레임에 있는 id 값으로 로그아웃 합니다.
// @Tags User
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer {AccessToken}"
// @Success 200
// @Failure 401 {object} models.ErrResponse
// @Failure 500 {object} models.ErrResponse
// @Router /users/logout [post]
func UserLogout(c *gin.Context) {
	au, err := auth.ExtractTokenMetadata(c.Request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, models.ErrResponse{
			Message: err.Error(),
		})
		return
	}

	if deleted, err := auth.DeleteAuth(au.AccessUuid); err != nil || deleted == 0 {
		c.JSON(http.StatusUnauthorized, models.ErrResponse{
			Message: err.Error(),
		})
		return
	}
	c.Status(http.StatusOK)
}

// @Summary 엑세스 토큰 검증
// @Description 올바르게 서명된 엑세스 토큰인지 검증합니다.
// @Tags User
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer {AccessToken}"
// @Success 200
// @Failure 401 {object} models.ErrResponse
// @Router /users/token/valid [get]
func UserTokenValid(c *gin.Context) {
	// 이미 미들웨어에서 서명 검사 및 redis에 존재여부를 모두 확인하고 호출하기 때문에 200만 응답해줍니다.
	c.Status(http.StatusOK)
}

// @Summary 리프레쉬 토큰 발급
// @Description 엑세스 토큰이 만료되었을때 리프레쉬 토큰을 이용하여 새롭게 발급하기 위하여 사용됩니다.
// @Tags User
// @Accept json
// @Produce json
// @Param data body models.JWTRefreshToken true "리프레쉬 토큰 정보"
// @Success 201 {object} models.JWTToken
// @Failure 401 {object} models.ErrResponse
// @Failure 403 {object} models.ErrResponse
// @Failure 422 {object} models.ErrResponse
// @Router /users/token/refresh [post]
func UserTokenRefresh(c *gin.Context) {
	refreshSecret := viper.GetString("JWT_REFRESH_SECRET")

	mapToken := models.JWTRefreshToken{}
	if err := c.ShouldBindJSON(&mapToken); err != nil {
		c.JSON(http.StatusUnprocessableEntity, models.ErrResponse{
			Message: err.Error(),
		})
		return
	}
	refreshToken := mapToken.RefreshToken

	//verify the token
	token, err := jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {
		//Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(refreshSecret), nil
	})
	//if there is an error, the token must have expired
	if err != nil {
		c.JSON(http.StatusUnauthorized, models.ErrResponse{
			Message: "Refresh token expired",
		})
		return
	}
	//is token valid?
	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		c.JSON(http.StatusUnauthorized, models.ErrResponse{
			Message: err.Error(),
		})
		return
	}
	//Since token is valid, get the uuid:
	claims, ok := token.Claims.(jwt.MapClaims) //the token claims should conform to MapClaims
	if ok && token.Valid {
		refreshUuid, ok := claims["refresh_uuid"].(string) //convert the interface to string
		if !ok {
			c.JSON(http.StatusUnprocessableEntity, models.ErrResponse{
				Message: err.Error(),
			})
			return
		}
		userId, err := strconv.ParseUint(fmt.Sprintf("%.f", claims["user_id"]), 10, 64)
		if err != nil {
			c.JSON(http.StatusUnprocessableEntity, models.ErrResponse{
				Message: "Error occurred",
			})
			return
		}
		//Delete the previous Refresh Token
		deleted, delErr := auth.DeleteAuth(refreshUuid)
		if delErr != nil || deleted == 0 {
			c.JSON(http.StatusUnauthorized, models.ErrResponse{
				Message: "unauthorized",
			})
			return
		}
		//Create new pairs of refresh and access tokens
		ts, createErr := auth.CreateToken(userId)
		if createErr != nil {
			c.JSON(http.StatusForbidden, models.ErrResponse{
				Message: createErr.Error(),
			})
			return
		}
		//save the tokens metadata to redis
		saveErr := auth.CreateAuth(userId, ts)
		if saveErr != nil {
			c.JSON(http.StatusForbidden, models.ErrResponse{
				Message: saveErr.Error(),
			})
			return
		}
		tokens := models.JWTToken{
			AccessToken:  ts.AccessToken,
			RefreshToken: ts.RefreshToken,
		}
		c.JSON(http.StatusCreated, tokens)
	} else {
		c.JSON(http.StatusUnauthorized, models.ErrResponse{
			Message: "refresh expired",
		})
	}
}

// @Summary 사용자 아바타 사진 업로드
// @Description 사용자 아바타(프로필) 사진을 업로드 합니다.
// @Tags User
// @Accept mpfd
// @Produce json
// @Param Authorization header string true "Bearer {AccessToken}"
// @Param file formData file true "아바타 사진"
// @Success 201
// @Failure 400 {object} models.ErrResponse
// @Failure 401 {object} models.ErrResponse
// @Failure 500 {object} models.ErrResponse
// @Router /users/avatar [post]
func UploadAvatar(c *gin.Context) {
	au, err := auth.ExtractTokenMetadata(c.Request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, models.ErrResponse{
			Message: err.Error(),
		})
		return
	}

	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrResponse{
			Message: err.Error(),
		})
		return
	}

	fileContent, err := file.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrResponse{
			Message: err.Error(),
		})
		return
	}
	defer fileContent.Close()

	fileData, err := ioutil.ReadAll(fileContent)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrResponse{
			Message: err.Error(),
		})
		return
	}

	if err := orm.Client.Model(&entitys.User{}).Where("id = ?", au.UserId).Update("AvatarImage", fileData).Error; err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrResponse{
			Message: err.Error(),
		})
		return
	}

	c.Status(http.StatusCreated)
}

// @Summary 사용자 아바타 사진 삭제
// @Description 사용자 아바타(프로필) 사진을 삭제합니다.
// @Tags User
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer {AccessToken}"
// @Success 204
// @Failure 401 {object} models.ErrResponse
// @Failure 500 {object} models.ErrResponse
// @Router /users/avatar [delete]
func DeleteAvatar(c *gin.Context) {
	au, err := auth.ExtractTokenMetadata(c.Request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, models.ErrResponse{
			Message: err.Error(),
		})
		return
	}

	if err := orm.Client.Model(&entitys.User{}).Where("id = ?", au.UserId).Update("AvatarImage", nil).Error; err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrResponse{
			Message: err.Error(),
		})
		return
	}

	c.Status(http.StatusNoContent)
}

// @Summary 사용자 삭제 (회원탈퇴)
// @Description 사용자를 삭제합니다.
// @Tags User
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer {AccessToken}"
// @Success 204
// @Failure 401 {object} models.ErrResponse
// @Failure 500 {object} models.ErrResponse
// @Router /users [delete]
func UserWithdrawal(c *gin.Context) {
	au, err := auth.ExtractTokenMetadata(c.Request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, models.ErrResponse{
			Message: err.Error(),
		})
		return
	}

	// jwt claim에는 user table의 id만 존재하여 카카오톡 소셜 테이블을 삭제하기 위한 정보를 얻기 위해 조회
	user := entitys.User{}
	if err := orm.Client.Where("id = ?", au.UserId).Find(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrResponse{
			Message: err.Error(),
		})
		return
	}

	// 카카오톡 소셜 테이블만 삭제하더라도 reference된 사용자 데이블도 cascade로 같이 삭제되지만 카카오 로그인이 아닌 회원가입의 경우 삭제되지 않기 때문에 2번의 쿼리로 사용자 데이터를 삭제하도록 합니다.
	if err := orm.Client.Where("id = ?", au.UserId).Delete(&entitys.User{}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrResponse{
			Message: err.Error(),
		})
		return
	}

	if err := orm.Client.Where("id = ?", user.KakaoTalkSocialsId).Delete(&entitys.KakaoTalkSocial{}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrResponse{
			Message: err.Error(),
		})
		return
	}

	if deleted, err := auth.DeleteAuth(au.AccessUuid); err != nil || deleted == 0 {
		c.JSON(http.StatusUnauthorized, models.ErrResponse{
			Message: err.Error(),
		})
		return
	}

	c.Status(http.StatusNoContent)
}
