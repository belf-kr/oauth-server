package entitys

import (
	"time"

	"github.com/parkgang/modern-board/internal/pkg/util"
	"gorm.io/gorm"
)

type User struct {
	Id                 uint            `gorm:"primary_key;auto_increment;" json:"id"`
	Email              string          `gorm:"unique;not null;size:50;" json:"email"`
	Password           string          `gorm:"size:64;" json:"password"`
	Name               string          `gorm:"not null;size:20;" json:"name"`
	AvatarImage        []byte          `json:"avatarImage"`
	ConnectedAt        time.Time       `gorm:"type:TIMESTAMP;default:current_timestamp;not null;" json:"connectedAt"`
	KakaoTalkSocialsId uint            `json:"kakaoTalkSocialsId"`
	KakaoTalkSocial    KakaoTalkSocial `gorm:"foreignkey:KakaoTalkSocialsId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"kakaoTalkSocial"`
}

func (u *User) BeforeSave(tx *gorm.DB) (err error) {
	// 비밀번호 암호화: sha2('{비밀번호}', 256)
	u.Password = util.Sha256(u.Password)

	// 로그인한 시간 UTC으로 저장: utc_timestamp()
	// 중요한 것은 GORM이 connection하는 loc에 의존됩니다: https://stackoverflow.com/a/60974094
	u.ConnectedAt = time.Now().UTC()
	return nil
}
