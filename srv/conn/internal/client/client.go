package client

import (
	"github.com/iobrother/zoo/core/config"
	"github.com/iobrother/zoo/core/log"
	"github.com/iobrother/zoo/core/transport/rpc/client"

	"github.com/getvox/vox/gen/rpc/channel"
	"github.com/getvox/vox/gen/rpc/chat"
	"github.com/getvox/vox/gen/rpc/sess"
)

var (
	chatClient    *chat.ChatClient
	channelClient *channel.ChannelClient
	sessClient    *sess.SessClient
)

type Registry struct {
	BasePath string
	EtcdAddr []string
}

func GetSessClient() *sess.SessClient {
	if sessClient == nil {
		r := &Registry{}
		if err := config.Scan("registry", &r); err != nil {
			log.Errorf("GetSessClient error=%v", err)
			return nil
		}
		cc, err := client.NewClient(
			client.WithServiceName("Sess"),
			client.BasePath(r.BasePath),
			client.EtcdAddr(r.EtcdAddr),
		)
		if err != nil {
			log.Errorf("GetSessClient error=%v", err)
			return nil
		}
		cli := cc.GetXClient()
		sessClient = sess.NewSessClient(cli)
	}
	return sessClient
}

func GetChatClient() *chat.ChatClient {
	if chatClient == nil {
		r := &Registry{}
		if err := config.Scan("registry", &r); err != nil {
			log.Errorf("GetChatClient error=%v", err)
			return nil
		}
		cc, err := client.NewClient(
			client.WithServiceName("Chat"),
			client.BasePath(r.BasePath),
			client.EtcdAddr(r.EtcdAddr),
		)
		if err != nil {
			log.Errorf("GetChatClient error=%v", err)
			return nil
		}
		cli := cc.GetXClient()
		chatClient = chat.NewChatClient(cli)
	}
	return chatClient
}

func GetChannelClient() *channel.ChannelClient {
	if channelClient == nil {
		r := &Registry{}
		if err := config.Scan("registry", &r); err != nil {
			log.Errorf("GetChannelClient error=%v", err)
			return nil
		}
		cc, err := client.NewClient(
			client.WithServiceName("Channel"),
			client.BasePath(r.BasePath),
			client.EtcdAddr(r.EtcdAddr),
		)
		if err != nil {
			log.Errorf("GetChannelClient error=%v", err)
			return nil
		}
		cli := cc.GetXClient()
		channelClient = channel.NewChannelClient(cli)
	}
	return channelClient
}
