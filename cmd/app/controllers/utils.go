package controllers

import (
	"fmt"

	"github.com/go-web-templates/api/internal/application/interfaces"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func bindAndValidate(ctx *fiber.Ctx, validator interfaces.JsonValidator, input interface{}) error {
	err := ctx.BodyParser(input)
	if err != nil {
		 return err
	}

	ok, errors := validator.Validate(input)
	if (!ok) {
		return ctx.Status(fiber.StatusBadRequest).JSON(errors)
	}

	return nil
}

func bindUUIDParam(ctx *fiber.Ctx, param string) (uuid.UUID, error) {
	rawID := ctx.Params(param)
	id , err := uuid.Parse(rawID)
	if (err != nil) {
		return id, ctx.Status(fiber.StatusBadRequest).SendString(
			fmt.Sprintf("invalid format for the %s parameter, expected an UUID", param),
		)
	}

	return id, nil
}
