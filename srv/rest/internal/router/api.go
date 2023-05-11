package router

import (
	"github.com/getvox/vox/gen/http/rest/channel"
	"github.com/getvox/vox/gen/http/rest/chat"
	sChannel "github.com/getvox/vox/srv/rest/internal/service/channel"
	sChat "github.com/getvox/vox/srv/rest/internal/service/chat"

	"github.com/iobrother/zoo/core/transport/http/server"
)

func RegisterAPI(s *server.Server) {
	g := s.Group("")

	chat.RegisterImHTTPService(g, sChat.GetService())
	channel.RegisterChannelHTTPService(g, sChannel.GetService())
}
