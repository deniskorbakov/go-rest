package event

import (
	"time"

	"github.com/go-list-templ/grpc/internal/domain/vo"
)

type UserEvent struct {
	UserID    vo.ID
	Event     TypeEvent
	EventTime time.Time
}
