package service

import (
	"context"
	"encoding/json"
	"time"

	"github.com/iobrother/zoo/core/log"
	"github.com/nats-io/nats.go"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/cast"
	"google.golang.org/protobuf/proto"

	"github.com/getvox/vox/gen/rpc/chat"
	"github.com/getvox/vox/gen/rpc/gid"
	"github.com/getvox/vox/pkg/constant"
	"github.com/getvox/vox/pkg/runtime"
	"github.com/getvox/vox/pkg/util"
	"github.com/getvox/vox/srv/chat/internal/client"
	"github.com/getvox/vox/srv/chat/internal/model"
)

type Chat struct{}

func GetChatService() *Chat {
	return &Chat{}
}

func (l *Chat) SendMsg(ctx context.Context, req *chat.SendReq, rsp *chat.SendRsp) (err error) {
	log.Infof("Chat SendMsg ChannelType=%d Type=%d Content=%s", req.ChannelType, req.MsgType, req.Content)
	now := time.Now().UnixMilli()
	m := chat.Msg{
		Id:            0,
		ChannelType:   req.ChannelType,
		Type:          req.MsgType,
		Content:       req.Content,
		From:          req.From,
		To:            req.To,
		SendTime:      now,
		Uuid:          req.Uuid,
		AtUserList:    req.AtUserList,
		IsTransparent: req.IsTransparent,
	}

	gidClient := client.GetGidClient()
	getRsp, err := gidClient.Get(ctx, &gid.GetReq{})
	if err != nil {
		return
	}
	m.Id = getRsp.Id

	b, err := proto.Marshal(&m)
	if err != nil {
		return
	}
	nm := &nats.Msg{
		Subject: "MSGS.new",
		Reply:   "",
		Data:    b,
		Sub:     nil,
	}
	js := runtime.GetJS()
	if _, err = js.PublishMsg(nm); err != nil {
		return
	}

	rsp.Id = m.Id
	rsp.SendTime = m.SendTime
	rsp.ClientUuid = m.Uuid

	return nil
}

// SyncMsg 同步离线消息，从redis缓存中读取，只同步最近7天的消息
func (l *Chat) SyncMsg(ctx context.Context, req *chat.SyncMsgReq, rsp *chat.SyncMsgRsp) (err error) {
	// 保证消息库 与 同步库 数据一致
	if err = l.removeDirty(ctx, req); err != nil {
		return
	}

	if req.Limit <= 0 {
		req.Limit = 20
	}
	if req.Limit > 100 {
		req.Limit = 100
	}

	zr := redis.ZRangeBy{
		Min:    "(" + cast.ToString(req.Offset),
		Max:    "+inf",
		Offset: 0,
		Count:  req.Limit,
	}

	rc := runtime.GetRedisClient()

	key := util.KeyMsgSync(req.Uin)
	cmd := rc.ZRangeByScore(ctx, key, &zr)
	val, err := cmd.Result()
	if err != nil {
		return
	}

	keys := make([]string, 0, len(val))
	for _, v := range val {
		key = util.KeyMsg(req.Uin, cast.ToInt64(v))
		keys = append(keys, key)
	}

	if len(keys) == 0 {
		return
	}
	rr, err := rc.MGet(ctx, keys...).Result()
	if err != nil {
		return err
	}
	for _, v := range rr {
		if v == nil {
			continue
		}
		msg := chat.Msg{}
		if err := json.Unmarshal([]byte(v.(string)), &msg); err != nil {
			continue
		}
		rsp.List = append(rsp.List, &msg)
	}

	return nil
}

func (l *Chat) removeDirty(ctx context.Context, req *chat.SyncMsgReq) (err error) {
	rc := runtime.GetRedisClient()
	// 删除过期id
	t := time.Now().AddDate(0, 0, -constant.MsgKeepDays)
	max := t.UnixNano() / 1e6
	key := util.KeyMsgSync(req.Uin)
	_, err = rc.ZRemRangeByScore(ctx, key, "-inf", cast.ToString(max)).Result()
	if err != nil {
		return
	}

	for {
		zr := redis.ZRangeBy{
			Min:    "(" + cast.ToString(req.Offset),
			Max:    "+inf",
			Offset: 0,
			Count:  1000,
		}

		key = util.KeyMsgSync(req.Uin)
		cmd := rc.ZRangeByScore(ctx, key, &zr)
		val, errr := cmd.Result()
		if errr != nil {
			return errr
		}

		var keys []string
		for _, v := range val {
			key = util.KeyMsg(req.Uin, cast.ToInt64(v))
			keys = append(keys, key)
		}
		if len(keys) == 0 {
			break
		}

		// 同步库中存在，而消息库中却不存在
		// 发生这种情况是因为，消息库中的消息过期已从redis中清除了，但是同步库中的消息id还未即时跑批处理清理掉
		var dirtyMembers []interface{}
		rr, errr := rc.MGet(ctx, keys...).Result()
		if errr != nil {
			return errr
		}
		for i, v := range rr {
			if v == nil {
				dirtyMembers = append(dirtyMembers, val[i])
				continue
			}
		}
		if len(dirtyMembers) == 0 {
			break
		} else {
			key = util.KeyMsgSync(req.Uin)
			if _, err = rc.ZRem(ctx, key, dirtyMembers...).Result(); err != nil {
				return
			}
		}
	}

	return nil
}

func (l *Chat) MsgAck(ctx context.Context, req *chat.MsgAckReq, rsp *chat.MsgAckRsp) (err error) {
	// TODO: 优化
	db := runtime.GetDB()
	msg := model.Msg{Id: req.Id}
	err = db.Take(&msg).Error
	if err != nil {
		return
	}
	rc := runtime.GetRedisClient()
	key := util.KeyMsgSync(req.Uin)
	rc.ZRemRangeByScore(ctx, key, "-inf", cast.ToString(msg.SendTime))
	return
}

func (l *Chat) DeleteMsg(ctx context.Context, req *chat.DeleteMsgReq, rsp *chat.DeleteMsgRsp) (err error) {
	if len(req.Ids) == 0 {
		return
	}

	rc := runtime.GetRedisClient()

	members := make([]any, 0, len(req.Ids))
	for _, id := range req.Ids {
		members = append(members, id)
	}
	key := util.KeyMsgSync(req.Uin)
	if err = rc.ZRem(context.Background(), key, members...).Err(); err != nil {
		return
	}
	return
}
