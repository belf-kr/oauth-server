package models

type KakaoToken struct {
	TokenType             string `json:"token_type" binding:"required"`               // 토큰 타입, bearer로 고정
	AccessToken           string `json:"access_token" binding:"required"`             // 사용자 액세스 토큰 값
	ExpiresIn             int    `json:"expires_in" binding:"required"`               // 액세스 토큰 만료 시간(초)
	RefreshToken          string `json:"refresh_token" binding:"required"`            // 사용자 리프레시 토큰 값
	RefreshTokenExpiresIn int    `json:"refresh_token_expires_in" binding:"required"` // 리프레시 토큰 만료 시간(초)
	Scope                 string `json:"scope"`                                       // 인증된 사용자의 정보 조회 권한 범위, 범위가 여러 개일 경우, 공백으로 구분
}
