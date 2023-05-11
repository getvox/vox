package client

import (
	"github.com/iobrother/zoo/core/config"
	"github.com/iobrother/zoo/core/transport/rpc/client"

	"github.com/iobrother/zim/gen/rpc/channel"
	"github.com/iobrother/zim/gen/rpc/chat"
)

var (
	chatClient    *chat.ChatClient
	channelClient *channel.ChannelClient
)

type Registry struct {
	BasePath string
	EtcdAddr []string
}

func GetChatClient() *chat.ChatClient {
	if chatClient == nil {
		r := &Registry{}
		config.Scan("registry", &r)

		// TODO: 优化
		cc, _ := client.NewClient(
			client.WithServiceName("Chat"),
			client.BasePath(r.BasePath),
			client.EtcdAddr(r.EtcdAddr),
		)
		cli := cc.GetXClient()
		chatClient = chat.NewChatClient(cli)
	}

	return chatClient
}

func GetChannelClient() *channel.ChannelClient {
	if channelClient == nil {
		r := &Registry{}
		config.Scan("registry", &r)

		// TODO: 优化
		cc, _ := client.NewClient(
			client.WithServiceName("Group"),
			client.BasePath(r.BasePath),
			client.EtcdAddr(r.EtcdAddr),
		)
		cli := cc.GetXClient()
		channelClient = channel.NewChannelClient(cli)
	}

	return channelClient
}
