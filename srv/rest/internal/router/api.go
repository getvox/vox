package router

import (
	"github.com/iobrother/zim/gen/http/rest/chat"
	"github.com/iobrother/zim/gen/http/rest/group"
	sChat "github.com/iobrother/zim/srv/rest/internal/service/chat"
	sGroup "github.com/iobrother/zim/srv/rest/internal/service/group"

	"github.com/iobrother/zoo/core/transport/http/server"
)

func RegisterAPI(s *server.Server) {
	g := s.Group("")

	chat.RegisterImHTTPService(g, sChat.GetService())
	group.RegisterGroupHTTPService(g, sGroup.GetService())
}
