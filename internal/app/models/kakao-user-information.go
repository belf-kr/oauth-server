package models

type KakaoUserInformation struct {
	ConnectedAt  string       `json:"connected_at" binding:"required"`
	Id           uint         `json:"id" binding:"required"`
	KakaoAccount kakaoAccount `json:"kakao_account"`
}

type kakaoAccount struct {
	Email   string  `json:"email"`
	Profile profile `json:"profile"`
}

type profile struct {
	Nickname          string `json:"nickname"`
	ProfileImageUrl   string `json:"profile_image_url"`
	ThumbnailImageUrl string `json:"thumbnail_image_url"`
}
