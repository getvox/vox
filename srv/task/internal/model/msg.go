package model

type Msg struct {
	Id         int64  `json:"id" gorm:"primaryKey;autoIncrement:false;comment:消息ID"`
	Type       int    `json:"type" gorm:"size:32;not null;default:0;comment:消息类型"`
	ConvType   int    `json:"conv_type" gorm:"size:8;not null;default:0;comment:会话类型"`
	Content    string `json:"content" gorm:"size:5000;not null;default:'';comment:消息内容"`
	Sender     string `json:"sender" gorm:"size:64;not null;default:'';comment:发送方"`
	Target     string `json:"target" gorm:"size:64;not null;default:'';comment:目标"`
	AtUserList string `json:"at_user_list" gorm:"size:1024;not null;default:'';comment:at列表"`
	SendTime   int64  `json:"send_time" gorm:"size:64;not null;default:0;comment:发送时间"`
	ReadTime   int64  `json:"read_time" gorm:"size:64;not null;default:0;comment:读时间"`
	ClientUuid string `json:"client_uuid" gorm:"size:64;not null;default:'';comment:消息指纹"`
}

func (_ *Msg) TableName() string {
	return "msg"
}
