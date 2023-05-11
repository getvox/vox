package service

import (
	"context"

	"github.com/iobrother/zoo/core/log"
	"github.com/spf13/cast"
	"gorm.io/gorm"

	"github.com/getvox/vox/gen/rpc/channel"
	"github.com/getvox/vox/gen/rpc/gid"
	"github.com/getvox/vox/pkg/runtime"
	"github.com/getvox/vox/srv/chat/internal/client"
	"github.com/getvox/vox/srv/chat/internal/model"
)

type Channel struct{}

func GetChannelService() *Channel {
	return &Channel{}
}

func (g *Channel) Create(ctx context.Context, req *channel.CreateReq, rsp *channel.CreateRsp) (err error) {
	ch := model.Channel{
		Owner: req.Owner,
		Cid:   "",
		Type:  0,
		Name:  req.Name,
	}

	gidClient := client.GetGidClient()
	getRsp, err := gidClient.Get(ctx, &gid.GetReq{})
	if err != nil {
		return
	}
	ch.Id = getRsp.Id

	if req.Cid != "" {
		ch.Cid = req.Cid
	} else {
		ch.Cid = cast.ToString(ch.Id)
	}

	getBatchReq := gid.GetBatchReq{Count: int32(len(req.Members) + 1)}
	getBatchRsp, err := gidClient.GetBatch(ctx, &getBatchReq)
	if err != nil {
		return
	}

	members := make([]*model.Member, 0, len(req.Members)+1)
	members = append(members, &model.Member{
		Id:     getBatchRsp.Ids[0],
		Cid:    ch.Cid,
		Member: req.Owner,
	})
	i := 1
	for _, v := range req.Members {
		if v == req.Owner {
			continue
		}
		member := &model.Member{
			Id:     getBatchRsp.Ids[i],
			Cid:    ch.Cid,
			Member: v,
		}
		members = append(members, member)
		i++
	}

	db := runtime.GetDB()
	err = db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&ch).Error; err != nil {
			return err
		}
		if err := tx.Create(&members).Error; err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		log.Error(err)
		return
	}
	rsp.Cid = ch.Cid

	return nil
}

func (g *Channel) GetJoinedList(ctx context.Context, req *channel.GetJoinedListReq, rsp *channel.GetJoinedListRsp) (err error) {
	db := runtime.GetDB()
	var rows []*model.Channel
	err = db.Model(&model.Member{}).Where(&model.Member{Member: req.Uin}).
		Select([]string{
			"channel.owner",
			"channel.cid",
			"channel.type",
			"channel.name",
			"channel.created_at",
			"channel.updated_at",
		}).
		Joins("INNER JOIN `channel` on member.cid=channel.cid").
		Find(&rows).Error

	for _, v := range rows {
		channelInfo := channel.ChannelInfo{
			Owner:     v.Owner,
			Name:      v.Name,
			Cid:       v.Cid,
			CreatedAt: v.CreatedAt.Unix(),
			UpdatedAt: v.UpdatedAt.Unix(),
			Type:      int32(v.Type),
		}
		rsp.List = append(rsp.List, &channelInfo)
	}

	return
}

func (g *Channel) Sync(ctx context.Context, req *channel.SyncReq, rsp *channel.SyncRsp) (err error) {
	if req.Limit == 0 {
		req.Limit = 20
	} else if req.Limit > 100 {
		req.Limit = 100
	}

	db := runtime.GetDB()
	err = db.Model(&model.Channel{}).Where(&model.Member{Member: req.Uin}).
		Select([]string{
			"channel.owner",
			"channel.cid",
			"channel.type",
			"channel.name",
			"UNIX_TIMESTAMP(channel.created_at) AS created_at",
			"UNIX_TIMESTAMP(channel.updated_at) AS updated_at",
		}).
		Scopes(func(db *gorm.DB) *gorm.DB {
			if req.Offset > 0 {
				db = db.Where("UNIX_TIMESTAMP(channel.created_at) > ?", req.Offset)
			}
			return db
		}).
		Joins("INNER JOIN channel on member.cid=channel.cid").
		Order("channel.updated_at ASC").
		Find(&rsp.List).Error

	return
}

func (g *Channel) Join(ctx context.Context, req *channel.JoinReq, rsp *channel.JoinRsp) (err error) {
	db := runtime.GetDB()
	v := model.Member{}
	err = db.Model(&model.Member{}).Find(&v, &model.Member{Cid: req.Cid, Member: req.Uin}).Error
	if v.Id == 0 {
		if err = db.Create(&model.Member{Cid: req.Cid, Member: req.Uin}).Error; err != nil {
			log.Error(err)
			return
		}
	}
	return
}

func (g *Channel) InviteUser(ctx context.Context, req *channel.InviteUserReq, rsp *channel.InviteUserRsp) (err error) {
	// TODO: 判断群是否存在
	gidClient := client.GetGidClient()
	getBatchReq := gid.GetBatchReq{Count: int32(len(req.UserList))}
	getBatchRsp, err := gidClient.GetBatch(ctx, &getBatchReq)
	if err != nil {
		return
	}

	db := runtime.GetDB()
	members := make([]*model.Member, 0, len(req.UserList))
	i := 0
	for _, v := range req.UserList {
		member := &model.Member{
			Id:     getBatchRsp.Ids[i],
			Cid:    req.Cid,
			Member: v,
		}
		members = append(members, member)
		i++
	}

	err = db.Create(&members).Error

	return
}

func (g *Channel) Quit(ctx context.Context, req *channel.QuitReq, rsp *channel.QuitRsp) (err error) {
	return
}

func (g *Channel) KickMember(ctx context.Context, req *channel.KickReq, rsp *channel.KickRsp) (err error) {
	return
}

func (g *Channel) Dismiss(ctx context.Context, req *channel.DismissReq, rsp *channel.DismissRsp) (err error) {
	return
}

func (g *Channel) GetMemberList(ctx context.Context, req *channel.GetMemberListReq, rsp *channel.GetMemberListRsp) (err error) {
	if req.Limit == 0 {
		req.Limit = 20
	} else if req.Limit > 100 {
		req.Limit = 100
	}

	db := runtime.GetDB()
	err = db.Model(&model.Member{}).Where(&model.Member{Cid: req.Cid}).
		Select([]string{
			"cid",
			"member",
			"UNIX_TIMESTAMP(created_at) AS created_at",
			"UNIX_TIMESTAMP(updated_at) AS updated_at",
		}).
		Scopes(func(db *gorm.DB) *gorm.DB {
			if req.Offset > 0 {
				db = db.Where("UNIX_TIMESTAMP(member.updated_at) > ?", req.Offset)
			}
			return db
		}).
		Order("updated_at ASC").
		Find(&rsp.List).Error

	return
}

func (g *Channel) GetMemberInfo(ctx context.Context, req *channel.GetMemberInfoReq, rsp *channel.GetMemberInfoRsp) (err error) {
	//db := runtime.GetDB()
	//v := model.Member{}
	//if err = db.Model(&model.Member{}).
	//	Find(&v, &model.Member{Cid: req.Cid, Member: req.Member}).
	//	Error; err != nil {
	//	log.Error(err)
	//	return
	//}
	//if v.Id != 0 {
	//	*rsp = channel.GetMemberInfoRsp{
	//		Cid:   v.Cid,
	//		Member:    v.Member,
	//		CreatedAt: v.CreatedAt.Unix(),
	//		UpdatedAt: v.UpdatedAt.Unix(),
	//	}
	//}
	return
}

func (g *Channel) SetMemberInfo(ctx context.Context, req *channel.SetMemberInfoReq, rsp *channel.SetMemberInfoRsp) (err error) {
	//db := runtime.GetDB()
	//if err = db.Model(&model.Member{}).
	//	Where(&model.Member{
	//		Cid: req.Cid,
	//		Member:  req.Member,
	//	}).
	//	Updates(&model.Member{
	//		Nickname: req.Nickname,
	//	}).Error; err != nil {
	//	log.Error(err)
	//	return
	//}
	return
}
