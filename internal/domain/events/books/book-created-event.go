package books

import (
	"time"

	"github.com/go-web-templates/api/internal/domain/entities"
)

type BookCreatedEvent struct {
	EventKind int           `json:"event_kind"`
	NewBook   entities.Book `json:"new_book"`
	Timestamp time.Time     `json:"timestamp"`
}

func NewBookCreatedEvent(newBook entities.Book) BookCreatedEvent {
	return BookCreatedEvent{
		EventKind: EVENT_KIND_CREATED,
		NewBook:   newBook,
		Timestamp: time.Now(),
	}
}

func (e BookCreatedEvent) GetBookEventKind() int {
	return e.EventKind
}
