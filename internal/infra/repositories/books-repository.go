package repositories

import (
	"database/sql"

	"github.com/go-web-templates/api/internal/domain/entities"
	"github.com/go-web-templates/api/internal/infra/data"
	"github.com/google/uuid"
)

type BooksRepository struct {
	db *data.Database
}

func NewBooksRepository(db *data.Database) *BooksRepository {
	return &BooksRepository{ 
		db,
	}
}

func (br *BooksRepository) Get(id uuid.UUID) (entities.Book, error) {
	model := entities.Book{}
	err := br.db.Ctx.QueryRow(
		`SELECT books.id, books.title, books.author, books.created_at
		 FROM books 
		 WHERE books.id = $1
		`,
		id,
	).Scan(&model.ID, &model.Title, &model.Author, &model.CreatedAt)

	return model, err
}

func (br *BooksRepository) GetAll() ([]entities.Book, error) {
	models := []entities.Book{}
	rows, err := br.db.Ctx.Query(
		`SELECT books.id, books.title, books.author, books.created_at
		 FROM books
		`,
	)

	if (err == sql.ErrNoRows) {
		return models, nil
	}

	for rows.Next() {
		model := entities.Book{}
		rows.Scan(&model.ID, &model.Title, &model.Author, &model.CreatedAt)
		models = append(models, model)
	}

	return models, err
}

func (br *BooksRepository) Create(entity *entities.Book) error {
	err := br.db.Ctx.QueryRow(`
		INSERT INTO books (id, title, author, created_at)
		VALUES ($1, $2, $3, $4)
		RETURNING id, title, author, created_at
		`,
		entity.ID, entity.Title, entity.Author, entity.CreatedAt,
	).Scan(&entity.ID, &entity.Title, &entity.Author, &entity.CreatedAt)

	return err
}

func (br *BooksRepository) Update(id uuid.UUID, entity *entities.Book) error {
	err := br.db.Ctx.QueryRow(
		`UPDATE books
		 SET title = $1, author = $2
		 WHERE id = $3
		 RETURNING id, title, author, created_at
		`,
		entity.Title, entity.Author, id,
	).Scan(&entity.ID, &entity.Title, &entity.Author, &entity.CreatedAt)

	return err
}

func (br *BooksRepository) Delete(id uuid.UUID) error {
	_, err := br.db.Ctx.Exec(
		`DELETE FROM books
		 WHERE id = $1
		`,
		id,
	)

	return err
}
