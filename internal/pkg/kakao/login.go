package kakao

import (
	"encoding/json"
	"log"
	"net/url"

	"github.com/belf-kr/oauth-server/internal/app/models"
	"github.com/belf-kr/oauth-server/internal/pkg/httpclient"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

// 카카오 사용자 정보를 가져옵니다.
func GetUserInfo(token string, v *models.KakaoUserInformation) error {
	endpoint := "https://kapi.kakao.com/v2/user/me"

	xWwwFormUrlencodedHelper := httpclient.NewXWwwFormUrlencodedHelperOptions()
	xWwwFormUrlencodedHelper.Header = map[string]string{
		"Authorization": "Bearer " + token,
	}

	resBody, err := xWwwFormUrlencodedHelper.Req(httpclient.POST, endpoint)
	if err != nil {
		return errors.Wrap(err, "kakao 사용자 정보 가져오기 실패")
	}
	log.Println("kakao 사용자 정보:" + string(resBody))

	if err := json.Unmarshal(resBody, &v); err != nil {
		return errors.Wrap(err, "kakao 사용자 정보 역직렬화 실패")
	}
	return nil
}

// 카카오 로그인이 완료되어 발급받은 인가 코드로 액세스 토큰과 리프레시 토큰을 발급 받습니다.
func GetToken(code string, v *models.KakaoToken) (err error) {
	endpoint := "https://kauth.kakao.com/oauth/token"
	restApiKey := viper.GetString("KAKAO_REST_API_KEY")
	redirectUri := viper.GetString("KAKAO_REDIRECT_URI")

	data := url.Values{}
	data.Set("grant_type", "authorization_code")
	data.Set("client_id", restApiKey)
	data.Set("redirect_uri", redirectUri)
	data.Set("code", code)

	xWwwFormUrlencodedHelper := httpclient.NewXWwwFormUrlencodedHelperOptions()
	xWwwFormUrlencodedHelper.Data = data
	resBody, err := xWwwFormUrlencodedHelper.Req(httpclient.POST, endpoint)
	if err != nil {
		return errors.Wrap(err, "kakao 토큰 받기 실패")
	}
	log.Println("카카오 로그인 후 REST API으로 토큰 받기 응답결과: " + string(resBody))

	// 응답 역직렬화
	if err := json.Unmarshal(resBody, &v); err != nil {
		return errors.Wrap(err, "kakao 발급받은 토큰 역직렬화 실패")
	}
	return nil
}
