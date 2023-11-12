package cacherepositories

import (
	"context"

	"github.com/go-web-templates/api/internal/domain/entities"
	"github.com/go-web-templates/api/internal/infra/data"
	"github.com/google/uuid"
)

type BooksCacheRepository struct {
	cache *data.Cache
}

const booksKeyPrefix = "books"

func booksKeyWithPrefix(id uuid.UUID) string {
	return addPrefix(booksKeyPrefix, id.String())
}

func NewBooksCacheRepository(cache *data.Cache) *BooksCacheRepository {
	return &BooksCacheRepository{
		cache,
	}
}

func (r *BooksCacheRepository) Set(entity entities.Book) error {
	ctx := context.Background()

	return r.cache.Ctx.Set(
		ctx,
		booksKeyWithPrefix(entity.ID),
		mustSerializeToJson(entity),
		0,
	).Err()
}

func (r *BooksCacheRepository) Get(id uuid.UUID) (entities.Book, error) {
	ctx := context.Background()

	entity := entities.Book{}

	raw, err := r.cache.Ctx.Get(ctx, booksKeyWithPrefix(id)).Result()

	if err != nil {
		return entities.Book{}, err
	}

	mustDeserializeFromJson(raw, &entity)

	return entity, nil
}

func (r *BooksCacheRepository) Delete(id uuid.UUID) error {
	ctx := context.Background()

	_, err := r.cache.Ctx.Del(ctx, booksKeyWithPrefix(id)).Result()

	return err
}
