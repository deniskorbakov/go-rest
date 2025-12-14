package repo

import (
	"context"

	"github.com/go-list-templ/grpc/internal/domain/entity"
	"github.com/go-list-templ/grpc/internal/domain/vo"
)

type User interface {
	Store(context.Context) (entity.User, error)
	Update(context.Context, entity.User) (entity.User, error)
	Delete(context.Context, entity.User) error
	GetById(context.Context, vo.ID) (entity.User, error)
	List(context.Context) ([]entity.User, error)
}
