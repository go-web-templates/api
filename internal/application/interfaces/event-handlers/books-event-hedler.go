package eventhandlers

import (
	events "github.com/go-web-templates/api/internal/domain/events/books"
)

type BooksEventHandler interface {
	Handle(event events.BookEvent) 
}
