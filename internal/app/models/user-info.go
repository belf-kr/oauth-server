package models

import (
	"encoding/base64"
	"net/http"

	"github.com/belf-kr/oauth-server/internal/app/entitys"
)

type UserInfo struct {
	Id          int    `json:"id" binding:"required" example:"1"`
	Email       string `json:"email" binding:"required" example:"user01@test.com"`
	Name        string `json:"name" binding:"required" example:"사용자01"`
	AvatarImage string `json:"avatarImage" binding:"required" example:"base64으로 인코딩된 이미지"`
}

func NewUserInfo(user entitys.User) UserInfo {
	// [base64 encoding for any image](https://freshman.tech/snippets/go/image-to-base64/)
	var base64Encoding string

	// 이미지 파일의 콘텐츠 유형에 맞게 적절한 URI 체계 헤더를 추가합니다.
	mimeType := http.DetectContentType(user.AvatarImage)
	switch mimeType {
	case "image/jpeg":
		base64Encoding += "data:image/jpeg;base64,"
	case "image/png":
		base64Encoding += "data:image/png;base64,"
	}

	base64Encoding += base64.StdEncoding.EncodeToString(user.AvatarImage)

	return UserInfo{
		Id:          int(user.Id),
		Email:       user.Email,
		Name:        user.Name,
		AvatarImage: base64Encoding,
	}
}
