package models

type UserSignUp struct {
	Name            string `json:"name" binding:"required" example:"사용자01"`
	Email           string `json:"email" binding:"required" example:"user01@test.com"`
	Password        string `json:"password" binding:"required" example:"test1!"`
	PasswordConfirm string `json:"passwordConfirm" binding:"required" example:"test1!"`
}
