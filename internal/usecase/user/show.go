package user

import (
	"context"

	"github.com/go-list-templ/grpc/internal/domain/entity"
	"github.com/go-list-templ/grpc/internal/domain/vo"
)

func (u *UseCase) Show(context.Context, vo.ID) (entity.User, error) {
	return entity.User{}, nil
}
