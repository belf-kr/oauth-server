package entitys

type KakaoTalkSocial struct {
	Id                uint   `gorm:"primary_key;" json:"id"`
	Email             string `gorm:"unique;not null;size:50;" json:"email"`
	NickName          string `gorm:"not null;size:20;" json:"nickName"`
	ProfileImageUrl   string `gorm:"size:2083;" json:"profileImageUrl"`   // 카카오톡에 실제로 저장된 프로필 사진 (고화질), 이미지의 url을 저장하기 위한 최대 크기로 잡아놓았습니다.
	ThumbnailImageUrl string `gorm:"size:2083;" json:"thumbnailImageUrl"` // 카카오톡에 저장된 프로필 사진 미리보기 (저화질), 이미지의 url을 저장하기 위한 최대 크기로 잡아놓았습니다.
}
