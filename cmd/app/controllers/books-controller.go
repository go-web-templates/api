package controllers

import (
	"github.com/go-web-templates/api/internal/application/dtos"
	"github.com/go-web-templates/api/internal/application/interfaces"
	"github.com/go-web-templates/api/internal/application/services"
	"github.com/gofiber/fiber/v2"
)

type BooksController struct{
	validator interfaces.JsonValidator
	service services.BooksService
}

func NewBooksController(
	validator interfaces.JsonValidator,
	service services.BooksService,
) *BooksController {
	return &BooksController{
		validator,
		service,
	}
}

func (bc *BooksController) GetAll(ctx *fiber.Ctx) error {
	books, err := bc.service.GetAll()
	if err != nil {
		return err
	}

	return ctx.JSON(books)
}

func (bc *BooksController) Get(ctx *fiber.Ctx) error {
	id, err := bindUUIDParam(ctx, "id")
	if (err != nil) {
		return err
	}

	book, err := bc.service.Get(id)
	if err != nil {
		return ctx.SendStatus(fiber.StatusNotFound)
	}

	return ctx.JSON(book)
}

func (bc *BooksController) Create(ctx *fiber.Ctx) error {
	input := new(dtos.BookInputDto)
	if err := bindAndValidate(ctx, bc.validator, input); err != nil {
		return err
	}

	book, err := bc.service.Create(*input)
	if (err != nil) {
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}
	
	return ctx.JSON(book)
}

func (bc *BooksController) Update(ctx *fiber.Ctx) error {
	id, err := bindUUIDParam(ctx, "id")
	if (err != nil) {
		return err
	}

	input := new(dtos.BookInputDto)
	if err := bindAndValidate(ctx, bc.validator, input); err != nil {
		return err
	}

	book, err := bc.service.Update(id, *input)
	if (err != nil) {
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}

	return ctx.JSON(book)
}

func (bc *BooksController) Delete(ctx *fiber.Ctx) error {
	id, err := bindUUIDParam(ctx, "id")
	if (err != nil) {
		return err
	}

	err = bc.service.Delete(id)
	if (err != nil) {
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}

	return ctx.SendStatus(fiber.StatusOK)
}

func (bc *BooksController) RegisterController(app *fiber.App) {
	router := app.Group("/books")

	router.Get("/", bc.GetAll)
	router.Get("/:id", bc.Get)
	router.Post("/", bc.Create)
	router.Put("/:id", bc.Update)
	router.Delete("/:id", bc.Delete)
}
