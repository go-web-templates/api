package entities

import (
	"time"

	"github.com/google/uuid"
)

type Book struct {
	ID        uuid.UUID `sql:"id" db:"books.id"`
	Title     string    `sql:"title" db:"books.title"`
	Author    string    `sql:"author" db:"books.author"`
	CreatedAt time.Time `sql:"created_at" db:"books.created_at"`
}
