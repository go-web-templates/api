package services

import (
	"github.com/go-web-templates/api/pkg/logger"
	"github.com/go-web-templates/api/internal/application/dtos"
	cacherepositories "github.com/go-web-templates/api/internal/application/interfaces/cache-repositories"
	eventhandlers "github.com/go-web-templates/api/internal/application/interfaces/event-handlers"
	"github.com/go-web-templates/api/internal/application/interfaces/repositories"
	"github.com/go-web-templates/api/internal/application/mappers"
	events "github.com/go-web-templates/api/internal/domain/events/books"
	"github.com/google/uuid"
)

type BooksService interface {
	GetAll() ([]dtos.BookOutputDto, error)
	Get(id uuid.UUID) (dtos.BookOutputDto, error)
	Create(input dtos.BookInputDto) (dtos.BookOutputDto, error)
	Update(id uuid.UUID, input dtos.BookInputDto) (dtos.BookOutputDto, error)
	Delete(id uuid.UUID) error
}

type ActualBooksService struct {
	logger          logger.ApplicationLogger
	repository      repositories.BooksRepository
	mapper          *mappers.BooksMapper
	cacheRepository cacherepositories.BooksCacheRepository
	eventsHandler   eventhandlers.BooksEventHandler
}

func NewActualBooksService(
	logger logger.ApplicationLogger,
	repository repositories.BooksRepository,
	mapper *mappers.BooksMapper,
	cacheRepository cacherepositories.BooksCacheRepository,
	eventsHandler eventhandlers.BooksEventHandler,
) *ActualBooksService {
	return &ActualBooksService{
		logger,
		repository,
		mapper,
		cacheRepository,
		eventsHandler,
	}
}

func (b *ActualBooksService) GetAll() ([]dtos.BookOutputDto, error) {
	entities, err := b.repository.GetAll()
	outputs := b.mapper.OutputsFromEntities(&entities)

	return outputs, err
}

func (b *ActualBooksService) Get(id uuid.UUID) (dtos.BookOutputDto, error) {
	entity, err := b.cacheRepository.Get(id)
	if err != nil {
		entity, err = b.repository.Get(id)

		if err == nil {
			b.logger.Warning(
				b.logger.Format("Cache miss for the book with ID: %s", id.String()),
				b.logger.Format("err: '%v'", err),
			)

			b.logger.Info(
				b.logger.Format("Adding %s to the cache", id.String()),
			)

			b.cacheRepository.Set(entity)
		}
	}

	output := b.mapper.OutputFromEntity(&entity)
	return output, err
}

func (b *ActualBooksService) Create(input dtos.BookInputDto) (dtos.BookOutputDto, error) {
	entity := b.mapper.EntityFromInput(&input)
	err := b.repository.Create(&entity)
	output := b.mapper.OutputFromEntity(&entity)

	event := events.NewBookCreatedEvent(entity)
	b.eventsHandler.Handle(event)

	return output, err
}

func (b *ActualBooksService) Update(id uuid.UUID, input dtos.BookInputDto) (dtos.BookOutputDto, error) {
	entity := b.mapper.EntityFromInput(&input)
	err := b.repository.Update(id, &entity)
	output := b.mapper.OutputFromEntity(&entity)

	event := events.NewBookUpdatedEvent(entity)
	b.eventsHandler.Handle(event)

	return output, err
}

func (b *ActualBooksService) Delete(id uuid.UUID) error {
	event := events.NewBookDeletedEvent(id)
	b.eventsHandler.Handle(event)

	return b.repository.Delete(id)
}
