package eventhandlers

import (
	"github.com/go-web-templates/api/pkg/logger"
	cacherepositories "github.com/go-web-templates/api/internal/application/interfaces/cache-repositories"
	events "github.com/go-web-templates/api/internal/domain/events/books"
)

type BooksEventHandler struct {
	cache cacherepositories.BooksCacheRepository
	logger logger.ApplicationLogger
}

func NewBooksEventHandler(
	cache cacherepositories.BooksCacheRepository,
	logger logger.ApplicationLogger,
) *BooksEventHandler {
	return &BooksEventHandler{
		cache,
		logger,
	}
}

// The ideal scenario is make this a async job with a queue
func (h *BooksEventHandler) Handle(event events.BookEvent) {
	switch e := event.(type) {

	case events.BookCreatedEvent:
		err := h.cache.Set(e.NewBook)
		h.logger.Info(
			h.logger.Format(
				"Creating the cache for the book with id: %s",
				e.NewBook.ID.String(),
			),
		)
		if (err != nil) {
			h.logger.Error(
				h.logger.Format(
					"Unable to create the cache for the book with id: %s",
					e.NewBook.ID.String(),
				),
				h.logger.Format("err: %v", err),
			)
		}

	case events.BookUpdatedEvent:
		err := h.cache.Set(e.NewBook)
		h.logger.Info(
			h.logger.Format(
				"Updating the cache for the book with id: %s",
				e.NewBook.ID.String(),
			),
		)
		if (err != nil) {
			h.logger.Error(
				h.logger.Format(
					"Unable to update the cache for the book with id: %s",
					e.NewBook.ID.String(),
				),
				h.logger.Format("err: %v", err),
			)
		}

	case events.BookDeletedEvent:
		err := h.cache.Delete(e.OldBookID)
		h.logger.Info(
			h.logger.Format(
				"Deleting the cache for the book with id: %s",
				e.OldBookID.String(),
			),
		)
		if (err != nil) {
			h.logger.Error(
				h.logger.Format(
					"Unable to delete the cache for the book with id: %s",
					e.OldBookID.String(),
				),
				h.logger.Format("err: %v", err),
			)
		}

	}
}
