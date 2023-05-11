package server

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/iobrother/zoo/core/log"
	"github.com/nats-io/nats.go"
	"github.com/redis/go-redis/v9"
	"google.golang.org/protobuf/proto"

	"github.com/iobrother/zim/gen/queue"
	"github.com/iobrother/zim/gen/rpc/sess"
	"github.com/iobrother/zim/pkg/constant"
	"github.com/iobrother/zim/pkg/runtime"
	"github.com/iobrother/zim/pkg/util"
	"github.com/iobrother/zim/srv/task/internal/client"
	"github.com/iobrother/zim/srv/task/internal/model"
)

type Server struct{}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Start() error {
	go s.consumeNew()
	go s.consumeTodo()

	log.Info("Dispatch Server Started")

	return nil
}

func (s *Server) Stop() error {
	return nil
}

func (s *Server) consumeTodo() {
	js := runtime.GetJS()
	sub, err := js.PullSubscribe("MSGS.todo", "TASK_TODO")
	if err != nil {
		log.Fatal(err)
	}

	for {
		msgs, err := sub.Fetch(10)
		if err != nil {
			if errors.Is(err, nats.ErrTimeout) {
				continue
			}
			log.Error(err.Error())
		} else {
			for _, m := range msgs {
				msg := queue.Msg{}
				if err := proto.Unmarshal(m.Data, &msg); err != nil {
					m.Ack()
					continue
				}

				if err := s.onTodo(&msg); err == nil {
					m.Ack()
				}
			}
		}
	}
}

func (s *Server) onTodo(m *queue.Msg) error {
	if err := s.storeRedis(m); err != nil {
		return err
	}

	s.push(m)

	return nil
}

func (s *Server) storeRedis(m *queue.Msg) error {
	if m.IsTransparent {
		return nil
	}

	member := redis.Z{
		Score:  float64(m.SendTime),
		Member: m.Id,
	}

	rc := runtime.GetRedisClient()
	b, err := json.Marshal(m)
	if err != nil {
		return err
	}
	// TODO: context
	ctx := context.Background()
	if _, err := rc.TxPipelined(ctx, func(pipe redis.Pipeliner) error {
		key := util.KeyMsgSync(m.InboxUser)
		pipe.ZAdd(ctx, key, member)
		pipe.Expire(ctx, key, time.Duration(constant.MsgKeepDays*24)*time.Hour)

		key = util.KeyMsg(m.InboxUser, m.Id)
		pipe.SetEx(ctx, key, string(b), time.Duration(constant.MsgKeepDays*24)*time.Hour)

		return nil
	}); err != nil {
		return err
	}
	return nil
}

func (s *Server) push(m *queue.Msg) {
	// 获取在线状态
	sessClient := client.GetSessClient()
	if sessClient != nil {
		log.Infof("Uin=%s", m.InboxUser)
		req := sess.GetOnlineReq{Uin: m.InboxUser}
		rsp, err := sessClient.GetOnline(context.Background(), &req)
		if err != nil {
			log.Error(err)
			return
		}

		m.InboxUser = ""
		b, err := proto.Marshal(m)
		if err != nil {
			log.Error(err)
			return
		}

		nc := runtime.GetNC()
		for _, v := range rsp.Conns {
			// online
			if v.Status == 1 {
				// 在线推送
				var onlines []string
				onlines = append(onlines, v.ConnId)
				pushMsg := queue.PushMsg{
					Server: v.Server,
					Conns:  onlines,
					Msg:    b,
				}
				bb, err := proto.Marshal(&pushMsg)
				if err != nil {
					log.Error(err)
					continue
				}

				mm := nats.Msg{
					Subject: fmt.Sprintf("push.online.%s", v.Server),
					Reply:   "",
					Header:  nil,
					Data:    bb,
					Sub:     nil,
				}
				if err := nc.PublishMsg(&mm); err != nil {
					log.Error(err)
				}
			} else if v.Status == 2 {
				// TODO: 离线推送
				log.Info("离线推送，待实现")
			}
		}
	} else {
		log.Info("client is null")
	}
}

func (s *Server) storeMysql(m *queue.Msg) {
	var atUserList string
	if len(m.AtUserList) > 0 {
		b, _ := json.Marshal(m.AtUserList)
		atUserList = string(b)
	}

	db := runtime.GetDB()
	msg := model.Msg{
		Id:          m.Id,
		ChannelType: int(m.ChannelType),
		Content:     m.Content,
		Type:        int(m.Type),
		From:        m.From,
		To:          m.To,
		AtUserList:  atUserList,
		ReadTime:    0,
		SendTime:    m.SendTime,
		Uuid:        m.Uuid,
	}

	if err := db.Create(&msg).Error; err != nil {
		log.Error(err)
	}
}

func (s *Server) consumeNew() {
	js := runtime.GetJS()
	sub, err := js.PullSubscribe("MSGS.new", "TASK_NEW")
	if err != nil {
		log.Fatal(err)
	}

	for {
		msgs, err := sub.Fetch(10)
		if err != nil {
			if errors.Is(err, nats.ErrTimeout) {
				continue
			}
			log.Error(err.Error())
		} else {
			for _, m := range msgs {
				msg := queue.Msg{}
				if err := proto.Unmarshal(m.Data, &msg); err != nil {
					log.Error(err)
					m.Ack()
					continue
				}

				if err := s.onNew(&msg); err == nil {
					m.Ack()
				}
			}
		}
	}
}

func (s *Server) onNew(m *queue.Msg) (err error) {
	if m.ChannelType == constant.ConvTypeC2C {
		err = s.onC2CMsg(m)
	} else if m.ChannelType == constant.ConvTypeGroup {
		err = s.onGroupMsg(m)
	}

	if err != nil {
		return
	}
	// 持久化，可以考虑生成一条 MSGS.persist，由独立进程做持久化
	go func() {
		s.storeMysql(m)
	}()

	return
}

func (s *Server) onC2CMsg(m *queue.Msg) error {
	js := runtime.GetJS()
	if m.From != "" {
		m.InboxUser = m.From
		b, err := proto.Marshal(m)
		if err != nil {
			return err
		}
		nm := &nats.Msg{
			Subject: "MSGS.todo",
			Reply:   "",
			Data:    b,
			Sub:     nil,
		}
		js.PublishMsg(nm)
	}

	if m.To != "" {
		m.InboxUser = m.To
		b, err := proto.Marshal(m)
		if err != nil {
			return err
		}
		nm := &nats.Msg{
			Subject: "MSGS.todo",
			Reply:   "",
			Data:    b,
			Sub:     nil,
		}
		js.PublishMsg(nm)
	}

	return nil
}

func (s *Server) onGroupMsg(m *queue.Msg) (err error) {
	db := runtime.GetDB()
	var members []*model.GroupMember
	cond := model.GroupMember{GroupId: m.To}
	if err = db.Where(&cond).Find(&members).Error; err != nil {
		return
	}

	js := runtime.GetJS()
	for _, v := range members {
		if v.Member == "" {
			continue
		}
		m.InboxUser = v.Member
		b, err := proto.Marshal(m)
		if err != nil {
			continue
		}
		nm := &nats.Msg{
			Subject: "MSGS.todo",
			Reply:   "",
			Data:    b,
			Sub:     nil,
		}

		js.PublishMsg(nm)
	}

	return
}
