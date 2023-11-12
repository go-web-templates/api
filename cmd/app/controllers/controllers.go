package controllers

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
)

var Module = fx.Provide(
	asController(NewIndexController),
	asController(NewBooksController),
)

type BaseController interface {
	RegisterController(app *fiber.App)
}

func asController(ctor any) any {
	return fx.Annotate(
		ctor,
		fx.As(new(BaseController)),
		fx.ResultTags(`group:"controllers"`),
	)
}

