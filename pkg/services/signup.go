package services

import (
	"context"
	"time"

	"github.com/ShariUllas/chat-app/pkg/models"
	"github.com/pkg/errors"
)

const addUsersQuery = `
INSERT INTO users
(
	name,
	email,
	password,
	created_at,
	updated_at
)
VALUES
(
	$1,$2,$3,$4,$5
)
returning id;
`

func (s *ChatApp) SignUp(ctx context.Context, request *models.SignUpRequest) (int, error) {
	var userID *int
	err := s.db.QueryRowContext(ctx, addUsersQuery, request.Name, request.Email, request.Password, time.Now(), time.Now()).Scan(&userID)
	if err != nil {
		return 0, errors.Wrapf(err, "adding users to db failed")
	}
	return *userID, nil
}
