package user

import (
	"context"

	"github.com/go-list-templ/grpc/internal/domain/entity"
)

func (u *UseCase) Update(context.Context, entity.User) (entity.User, error) {
	return entity.User{}, nil
}
