package service

import (
	"context"
	"encoding/json"
	"fmt"
	"sort"
	"time"

	"github.com/iobrother/zoo/core/log"
	"github.com/redis/go-redis/v9"

	"github.com/iobrother/zim/pkg/constant"
	"github.com/iobrother/zim/pkg/runtime"
	"github.com/iobrother/zim/pkg/util"
)

// 给不同类型设备增加不同的 TAG 标记
// 当设备冲突时，可以有两种策略，1：踢掉较早登录的设备（从登录页主动登录情况） 2：提示当前设备登录冲突（重连情况）

// 离线推送通知
// 推送服务的设备数据与用户 UIN 进行关联
// 云端自动将即时通讯消息转成特定的推送通知发送至客户端
// 根据 UIN 找到关联设备

// 离线消息同步
// 云端主动推送方案
// 客户端从云端拉取方案，用户登录上线后，计算出用户离线期间产生的未读消息的对话列表及对应的未读消息数，以未读消息更新事件通知到客户端

type ConnInfo struct {
	ConnID         string `json:"conn_id"`
	DeviceId       string `json:"device_id"`
	DeviceName     string `json:"device_name"`
	Tag            string `json:"tag"`
	Platform       string `json:"platform"`
	Server         string `json:"server"`
	LoginTime      int64  `json:"login_time"`
	DisconnectTime int64  `json:"disconnect_time"`
	Status         int    `json:"status"` // 获取状态，GetRealStatus()方法
}

func (d *ConnInfo) GetRealStatus() int {
	status := d.Status
	if d.DisconnectTime != 0 && d.Status == constant.PushOnline {
		if time.Since(time.Unix(d.DisconnectTime, 0)) > time.Duration(constant.PushOnlineKeepDays*24)*time.Hour {
			d.Status = constant.Offline
		}
	}
	return status
}

func (s *Service) addConn(ctx context.Context, uin string, info *ConnInfo) (err error) {
	b, err := json.Marshal(info)
	if err != nil {
		return
	}
	rc := runtime.GetRedisClient()
	_, err = rc.TxPipelined(ctx, func(pipe redis.Pipeliner) error {
		key := util.KeyOnline(uin, info.ConnID)
		pipe.Set(ctx, key, string(b), time.Minute*2)
		key = util.KeyDevice(uin)
		pipe.HSet(ctx, key, info.ConnID, string(b))
		return nil
	})
	return
}

func (s *Service) delConn(ctx context.Context, uin string, info *ConnInfo) (err error) {
	rc := runtime.GetRedisClient()
	_, err = rc.Pipelined(ctx, func(pipe redis.Pipeliner) error {
		pipe.Del(ctx, util.KeyOnline(uin, info.ConnID))
		return nil
	})

	return
}

func (s *Service) getConn(ctx context.Context, uin, id string) *ConnInfo {
	key := util.KeyDevice(uin)

	rc := runtime.GetRedisClient()
	if b, err := rc.HGet(ctx, key, id).Bytes(); err != nil {
		return nil
	} else {
		info := &ConnInfo{}
		if err := json.Unmarshal(b, info); err != nil {
			return nil
		}
		return info
	}
}

func (s *Service) getOnline(ctx context.Context, uin string) (onlines map[string][]*ConnInfo, err error) {
	onlines = make(map[string][]*ConnInfo)
	rc := runtime.GetRedisClient()
	keys, err := rc.Keys(ctx, fmt.Sprintf("online:%s:*", uin)).Result()
	if err != nil {
		return
	}
	if len(keys) == 0 {
		return
	}
	log.Infof("online keys=%v", keys)
	result, err := rc.MGet(ctx, keys...).Result()
	if err != nil {
		return
	}

	for _, v := range result {
		info := ConnInfo{}
		if err := json.Unmarshal([]byte(v.(string)), &info); err != nil {
			continue
		}
		onlines[info.Server] = append(onlines[info.Server], &info)
	}

	return
}

func (s *Service) getOnlineOfTag(ctx context.Context, uin string, tag string) (onlines []*ConnInfo, err error) {
	onlines = make([]*ConnInfo, 0)
	rc := runtime.GetRedisClient()
	keys, err := rc.Keys(ctx, fmt.Sprintf("online:%s:*", uin)).Result()
	if err != nil {
		return
	}
	if len(keys) == 0 {
		return
	}

	log.Infof("online keys=%v", keys)
	result, err := rc.MGet(ctx, keys...).Result()
	if err != nil {
		return
	}
	for _, v := range result {
		info := ConnInfo{}
		if err := json.Unmarshal([]byte(v.(string)), &info); err != nil {
			continue
		}

		if info.Status != constant.Online {
			continue
		}
		if info.Tag == tag {
			onlines = append(onlines, &info)
			return
		}
	}

	if len(onlines) > 1 {
		sort.Slice(onlines, func(i, j int) bool { return onlines[i].LoginTime < onlines[j].LoginTime })
	}
	return onlines, nil
}
