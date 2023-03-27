package router

import (
	"github.com/iobrother/zim/gen/http/rest/chat"

	sChat "github.com/iobrother/zim/srv/rest/internal/service/chat"
	"github.com/iobrother/zoo/core/transport/http/server"
)

func RegisterAPI(s *server.Server) {
	g := s.Group("")

	chat.RegisterImHTTPService(g, sChat.GetService())
}
