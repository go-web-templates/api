package books

import (
	"time"

	"github.com/go-web-templates/api/internal/domain/entities"
)

type BookUpdatedEvent struct {
	EventKind int           `json:"event_kind"`
	NewBook   entities.Book `json:"new_book"`
	Timestamp time.Time     `json:"timestamp"`
}

func NewBookUpdatedEvent(newBook entities.Book) BookUpdatedEvent {
	return BookUpdatedEvent{
		EventKind: EVENT_KIND_UPDATED,
		NewBook:   newBook,
		Timestamp: time.Now(),
	}
}

func (e BookUpdatedEvent) GetBookEventKind() int {
	return e.EventKind
}
