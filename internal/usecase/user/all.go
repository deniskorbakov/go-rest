package user

import (
	"context"

	"github.com/go-list-templ/grpc/internal/domain/entity"
)

func (u *UseCase) All(context.Context) ([]entity.User, error) {
	return []entity.User{}, nil
}
