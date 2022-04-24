package services

import (
	"context"

	"github.com/ShariUllas/chat-app/pkg/models"
	"github.com/pkg/errors"
)

type Users struct {
	ID       int
	Name     string
	Email    string
	Password string
}

const getUsersByEmailQuery = `
SELECT id,
       name,
	   email,
	   password
FROM   users
WHERE  email = $1
`

func (s *ChatApp) Login(ctx context.Context, request *models.LoginRequest) (Users, error) {
	var users Users
	if err := s.db.QueryRowContext(ctx, getUsersByEmailQuery, request.Email).Scan(
		&users.ID,
		&users.Name,
		&users.Email,
		&users.Password,
	); err != nil {
		return Users{}, errors.Wrap(err, "get category by id query scan failed")
	}
	if request.Password != users.Password {
		return Users{}, errors.New("password mismatch")
	}
	return users, nil
}
