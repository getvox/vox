package model

import (
	"time"
)

type Group struct {
	Id        int64     `json:"id" gorm:"primaryKey;autoIncrement;comment:系统编号"`
	Owner     string    `json:"owner" gorm:"size:64;not null;default:'';comment:群主"`
	GroupId   string    `json:"group_id" gorm:"size:64;not null;default:'';comment:群ID"`
	Type      int       `json:"type" gorm:"size:8;not null;default:0;comment:群类型"`
	Name      string    `json:"name" gorm:"size:64;not null;default:'';comment:群名称"`
	CreatedAt time.Time `json:"created_at" gorm:"comment:创建时间"`
	UpdatedAt time.Time `json:"updated_at" gorm:"comment:更新时间"`
}

func (_ *Group) TableName() string {
	return "group"
}
