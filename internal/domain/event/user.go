package event

import (
	"time"

	"github.com/go-list-templ/grpc/internal/domain/vo"
)

const (
	Created TypeUserEvent = "created"
	Deleted TypeUserEvent = "deleted"
)

type (
	TypeUserEvent string

	UserEvent struct {
		Event     TypeUserEvent
		UserID    vo.ID
		EventTime time.Time
	}
)
