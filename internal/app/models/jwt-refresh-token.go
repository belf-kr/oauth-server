package models

type JWTRefreshToken struct {
	RefreshToken string `json:"refreshToken" binding:"required" example:"header.payLoad.signature"`
}
