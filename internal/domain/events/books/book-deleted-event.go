package books

import (
	"time"

	"github.com/google/uuid"
)

type BookDeletedEvent struct {
	EventKind int       `json:"event_kind"`
	OldBookID uuid.UUID `json:"old_book_id"`
	Timestamp time.Time `json:"timestamp"`
}

func NewBookDeletedEvent(oldBookID uuid.UUID) BookDeletedEvent {
	return BookDeletedEvent{
		EventKind: EVENT_KIND_DELETED,
		OldBookID: oldBookID,
		Timestamp: time.Now(),
	}
}

func (e BookDeletedEvent) GetBookEventKind() int {
	return e.EventKind
}
