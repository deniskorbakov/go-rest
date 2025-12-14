package usecase

import (
	"context"

	"github.com/go-list-templ/grpc/internal/domain/entity"
	"github.com/go-list-templ/grpc/internal/domain/vo"
)

type (
	User interface {
		Create(context.Context) (entity.User, error)
		Update(context.Context, entity.User) (entity.User, error)
		Delete(context.Context, entity.User) error
		Show(context.Context, vo.ID) (entity.User, error)
		All(context.Context) ([]entity.User, error)
	}
)
