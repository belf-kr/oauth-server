package models

type UserInfo struct {
	Id          int    `json:"id" binding:"required" example:"1"`
	Email       string `json:"email" binding:"required" example:"user01@test.com"`
	Name        string `json:"name" binding:"required" example:"사용자01"`
	AvatarImage string `json:"avatarImage" binding:"required" example:"base64으로 인코딩된 이미지"`
}
