package event

import "github.com/go-list-templ/grpc/internal/domain/vo"

type UserCreated struct {
	UserID    vo.ID
	CreatedAt vo.Time
}

type UserDeleted struct {
	UserID    vo.ID
	DeletedAt vo.Time
}
