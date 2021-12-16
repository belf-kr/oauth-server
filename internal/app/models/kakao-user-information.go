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

// 프로필 사진의 경우 kakao_account.profile 과 properties 2곳에서 조회가 가능한데 가장 fresh한 데이터는 kakao_account.profile 에 존재합니다. properties 에서 프로필 사진 사용시 이전 프로필 사진이 사용될 수 있습니다.
type profile struct {
	Nickname          string `json:"nickname"`
	ProfileImageUrl   string `json:"profile_image_url"`
	ThumbnailImageUrl string `json:"thumbnail_image_url"`
}
