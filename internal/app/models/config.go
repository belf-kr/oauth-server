package models

type Config struct {
	Kakao Kakao `json:"kakao"`
}

type Kakao struct {
	RestApiKey  string `json:"restApiKey" example:"111aaaa1a111aaa1a111a1a1a11a1a11"`
	RedirectUri string `json:"redirectUri" example:"http://localhost:8080/api/users/login/kakao"`
}
