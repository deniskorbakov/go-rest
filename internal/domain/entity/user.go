package entity

import (
	"time"

	"github.com/go-list-templ/grpc/internal/domain/vo"
)

type User struct {
	ID        vo.ID
	Username  vo.Username
	Email     vo.Email
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewUser(username, email string) (*User, error) {
	validUsername, err := vo.NewUsername(username)
	if err != nil {
		return nil, err
	}

	validEmail, err := vo.NewEmail(email)
	if err != nil {
		return nil, err
	}

	return &User{
		ID:        vo.NewID(),
		Username:  validUsername,
		Email:     validEmail,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}, nil
}
