package services

import (
	"github.com/ShariUllas/chat-app/pkg/config"
)

type ChatApp struct {
	cfg *config.Config
}

func New(cfg *config.Config) *ChatApp {
	return &ChatApp{
		cfg: cfg,
	}
}

type ChatAppService interface {
}
