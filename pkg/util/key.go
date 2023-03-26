package util

import "fmt"

const (
	prefixDevice = "device:%s"
	prefixOnline = "online:%s:%s"

	prefixMsg     = "msg:%s:%d"   // msg:uin:id
	prefixMsgSync = "msg_sync:%s" // msg_sync:uin
)

func KeyDevice(uin string) string {
	return fmt.Sprintf(prefixDevice, uin)
}

func KeyOnline(uin, deviceId string) string {
	return fmt.Sprintf(prefixOnline, uin, deviceId)
}

func KeyMsg(uin string, id int64) string {
	return fmt.Sprintf(prefixMsg, uin, id)
}

func KeyMsgSync(uin string) string {
	return fmt.Sprintf(prefixMsgSync, uin)
}
