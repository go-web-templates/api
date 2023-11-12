package dtos

type BookInputDto struct {
	Title  string `json:"title" validate:"required,min=3,max=200"`
	Author string `json:"author" validate:"required,min=3,max=100"`
}
