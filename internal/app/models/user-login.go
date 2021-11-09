package models

type UserLogin struct {
	Email    string `json:"email" binding:"required" example:"user01@test.com"`
	Password string `json:"password" binding:"required" example:"test1!"`
}
