package model

import (
	"time"
)

type Member struct {
	Id        int64     `json:"id" gorm:"primaryKey;autoIncrement;comment:系统编号"`
	Cid       string    `json:"cid" gorm:"size:64;not null;default:'';comment:群ID"`
	Member    string    `json:"member" gorm:"size:64;not null;default:'';comment:群成员"`
	CreatedAt time.Time `json:"created_at" gorm:"comment:创建时间"`
	UpdatedAt time.Time `json:"updated_at" gorm:"comment:更新时间"`
}

func (_ *Member) TableName() string {
	return "member"
}
