package handlers

import (
	"net"

	"github.com/ShariUllas/chat-app/pkg/config"
	"github.com/ShariUllas/chat-app/pkg/services"

	"github.com/gin-gonic/gin"
)

type REST struct {
	Engine *gin.Engine
	Routes *gin.RouterGroup
	config *config.Config
}

func New(c *config.Config, service services.ChatAppService) *REST {
	engine := gin.New()

	routesGrp := engine.Group("")

	engine.Use(
		gin.Recovery(),
	)

	v1ChatAppGroup := routesGrp.Group("v1")

	routers := RouterGroups{
		ChatAppGroup: v1ChatAppGroup,
	}

	NewChatAppController(service, c, routers)

	return &REST{
		Engine: engine,
		Routes: routesGrp,
		config: c,
	}
}

func (r *REST) Run() {
	r.Engine.Run(net.JoinHostPort(r.config.Host, r.config.APIPort))
}
