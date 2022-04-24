package handlers

import (
	"github.com/ShariUllas/chat-app/pkg/config"
	"github.com/ShariUllas/chat-app/pkg/services"

	"github.com/gin-gonic/gin"
)

type ChatAppController struct {
	config  *config.Config
	service services.ChatAppService
}

type RouterGroups struct {
	ChatAppGroup *gin.RouterGroup
}

func NewChatAppController(service services.ChatAppService, cfg *config.Config, router RouterGroups) {
	c := &ChatAppController{
		config:  cfg,
		service: service,
	}
	router.ChatAppGroup.POST("/login", c.Login)
	router.ChatAppGroup.POST("/signup", c.SignUp)
}
