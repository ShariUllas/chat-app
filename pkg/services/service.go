package services

import (
	"context"
	"database/sql"

	"github.com/ShariUllas/chat-app/pkg/config"
	"github.com/ShariUllas/chat-app/pkg/models"
)

type ChatApp struct {
	cfg *config.Config
	db  *sql.DB
}

func New(cfg *config.Config, db *sql.DB) *ChatApp {
	return &ChatApp{
		cfg: cfg,
		db:  db,
	}
}

type ChatAppService interface {
	Login(ctx context.Context, request *models.LoginRequest) (Users, error)
	SignUp(ctx context.Context, request *models.SignUpRequest) (int, error)
}
