package dtos

import (
	"time"

	"github.com/google/uuid"
)

type BookOutputDto struct {
	ID        uuid.UUID `json:"id"`
	Title     string    `json:"title"`
	Author    string    `json:"author"`
	CreatedAt time.Time `json:"created_at"`
}
