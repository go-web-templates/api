package mappers

import (
	"time"

	"github.com/go-web-templates/api/internal/application/dtos"
	"github.com/go-web-templates/api/internal/domain/entities"
	"github.com/google/uuid"
)

type BooksMapper struct {}

func NewBooksMapper() *BooksMapper {
	return &BooksMapper{}
}

func (bm *BooksMapper) OutputFromEntity(entity *entities.Book) dtos.BookOutputDto {
	return dtos.BookOutputDto{
		ID: entity.ID,
		Title: entity.Title,
		Author: entity.Author,
		CreatedAt: entity.CreatedAt,
	}
}

func (bm *BooksMapper) OutputsFromEntities(entities *[]entities.Book) []dtos.BookOutputDto {
	outputs := make([]dtos.BookOutputDto, len(*entities))
	for i, entity := range *entities {
		outputs[i] = bm.OutputFromEntity(&entity)
	}
	return outputs
}

func (bm *BooksMapper) EntityFromInput(input *dtos.BookInputDto) entities.Book {
	return entities.Book{
		ID: uuid.New(),
		Title: input.Title,
		Author: input.Author,
		CreatedAt: time.Now(),
	}
}
