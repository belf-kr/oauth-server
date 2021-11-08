package models

type JWTToken struct {
	AccessToken  string `json:"accessToken" binding:"required" example:"header.payLoad.signature"`
	RefreshToken string `json:"refreshToken" binding:"required" example:"header.payLoad.signature"`
}
