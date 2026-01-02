package usecase

import (
	"context"

	"github.com/go-list-templ/grpc/internal/domain/entity"
)

type (
	User interface {
		Create(context.Context, entity.User) (entity.User, error)
		All(context.Context) ([]entity.User, error)
	}
)
