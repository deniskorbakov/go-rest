package user

import (
	"context"

	"github.com/go-list-templ/grpc/internal/domain/entity"
)

func (u *UseCase) Delete(context.Context, entity.User) error {
	return nil
}
