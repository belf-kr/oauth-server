package auth

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/belf-kr/oauth-server/internal/app/data/redis"
	"github.com/golang-jwt/jwt"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"github.com/twinj/uuid"
)

type TokenDetails struct {
	AccessToken  string
	RefreshToken string
	AccessUuid   string
	RefreshUuid  string
	AtExpires    int64
	RtExpires    int64
}

// AccessDetails Redis에서 조회해야 하는 메타데이터
type AccessDetails struct {
	AccessUuid string
	UserId     uint64
}

// jwt 토큰을 생성합니다.
func CreateToken(userid uint64) (*TokenDetails, error) {
	accessSecret := viper.GetString("ACCESS_SECRET")
	refreshSecret := viper.GetString("REFRESH_SECRET")

	td := &TokenDetails{}
	td.AtExpires = time.Now().Add(time.Minute * 15).Unix()
	td.AccessUuid = uuid.NewV4().String()

	td.RtExpires = time.Now().Add(time.Hour * 24 * 7).Unix()
	td.RefreshUuid = uuid.NewV4().String()

	var err error
	//Creating Access Token
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["access_uuid"] = td.AccessUuid
	atClaims["user_id"] = userid
	atClaims["exp"] = td.AtExpires
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	td.AccessToken, err = at.SignedString([]byte(accessSecret))
	if err != nil {
		return nil, err
	}
	//Creating Refresh Token
	rtClaims := jwt.MapClaims{}
	rtClaims["refresh_uuid"] = td.RefreshUuid
	rtClaims["user_id"] = userid
	rtClaims["exp"] = td.RtExpires
	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)
	td.RefreshToken, err = rt.SignedString([]byte(refreshSecret))
	if err != nil {
		return nil, err
	}
	return td, nil
}

// 토큰을 redis에 저장합니다.
func CreateAuth(userid uint64, td *TokenDetails) error {
	// converting Unix to UTC
	at := time.Unix(td.AtExpires, 0)
	rt := time.Unix(td.RtExpires, 0)
	now := time.Now()

	ctx := context.Background()

	errAccess := redis.Client.Set(ctx, td.AccessUuid, strconv.Itoa(int(userid)), at.Sub(now)).Err()
	if errAccess != nil {
		return errAccess
	}
	errRefresh := redis.Client.Set(ctx, td.RefreshUuid, strconv.Itoa(int(userid)), rt.Sub(now)).Err()
	if errRefresh != nil {
		return errRefresh
	}
	return nil
}

// TokenValid 토큰 만료 여부를 검사합니다.
func TokenValid(r *http.Request) error {
	token, err := VerifyToken(r)
	if err != nil {
		return err
	}
	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		return err
	}
	return nil
}

// VerifyToken 토큰 검증 합니다.
func VerifyToken(r *http.Request) (*jwt.Token, error) {
	accessSecret := viper.GetString("ACCESS_SECRET")
	tokenString := ExtractToken(r)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		//Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(accessSecret), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

// ExtractToken 요청 헤더(Request header)에서 토큰을 가져옵니다.
func ExtractToken(r *http.Request) string {
	bearToken := r.Header.Get("Authorization")
	//normally Authorization the_token_xxx
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}

// ExtractTokenMetadata Redis 저장소에서 토큰을 추출합니다.
func ExtractTokenMetadata(r *http.Request) (*AccessDetails, error) {
	token, err := VerifyToken(r)
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		accessUuid, ok := claims["access_uuid"].(string)
		if !ok {
			return nil, err
		}
		userId, err := strconv.ParseUint(fmt.Sprintf("%.f", claims["user_id"]), 10, 64)
		if err != nil {
			return nil, err
		}
		return &AccessDetails{
			AccessUuid: accessUuid,
			UserId:     userId,
		}, nil
	}
	return nil, err
}

// FetchAuth 토큰에 저장된 uuid를 Redis에서 찾는 로직
func FetchAuth(authD *AccessDetails) (uint64, error) {
	ctx := context.Background()

	userid, err := redis.Client.Get(ctx, authD.AccessUuid).Result()
	if err != nil {
		return 0, err
	}
	userID, _ := strconv.ParseUint(userid, 10, 64)
	return userID, nil
}

// 매개 변수로 전달된 UUID에 해당하는 Redis의 레코드를 삭제합니다.
func DeleteAuth(givenUuid string) (int64, error) {
	ctx := context.Background()

	deleted, err := redis.Client.Del(ctx, givenUuid).Result()
	if err != nil {
		return 0, err
	}
	return deleted, nil
}

// jwt token 발급과 동시에 redis에 세션을 저장합니다.
func CreateSession(userid uint64) (*TokenDetails, error) {
	ts, err := CreateToken(userid)
	if err != nil {
		return nil, errors.Wrap(err, "jwt token 생성에 실패하였습니다")
	}

	if err := CreateAuth(userid, ts); err != nil {
		return nil, errors.Wrap(err, "token을 redis에 저장하지 못하였습니다")
	}

	return ts, nil
}
