package model

type User struct {
	Id       int64  `json:"id" gorm:"primaryKey;autoIncrement;comment:系统编号"`
	Uid      string `json:"uid" gorm:"size:64;not null;default:'';comment:用户识别号"`
	Nickname string `json:"nickname" gorm:"size:64;not null;default:'';comment:昵称"`
}

func (_ *User) TableName() string {
	return "user"
}
