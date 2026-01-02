package repo

import (
	"context"

	"github.com/go-list-templ/grpc/internal/domain/entity"
)

type (
	UserPersistentRepo interface {
		Store(context.Context, entity.User) error
		All(context.Context) ([]entity.User, error)
	}
)
