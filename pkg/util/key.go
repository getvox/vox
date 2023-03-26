package util

import "fmt"

const (
	prefixDevice = "device:%s"
	prefixOnline = "online:%s:%s"
)

func KeyDevice(uin string) string {
	return fmt.Sprintf(prefixDevice, uin)
}

func KeyOnline(uin, deviceId string) string {
	return fmt.Sprintf(prefixOnline, uin, deviceId)
}
