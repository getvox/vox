package client

import (
	"github.com/iobrother/zoo/core/config"
	"github.com/iobrother/zoo/core/transport/rpc/client"

	"github.com/iobrother/zim/gen/rpc/chat"
	"github.com/iobrother/zim/gen/rpc/group"
)

var (
	chatClient  *chat.ChatClient
	groupClient *group.GroupClient
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

func GetGroupClient() *group.GroupClient {
	if groupClient == nil {
		r := &Registry{}
		config.Scan("registry", &r)

		// TODO: 优化
		cc, _ := client.NewClient(
			client.WithServiceName("Group"),
			client.BasePath(r.BasePath),
			client.EtcdAddr(r.EtcdAddr),
		)
		cli := cc.GetXClient()
		groupClient = group.NewGroupClient(cli)
	}

	return groupClient
}
