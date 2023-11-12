package services

import (
	"github.com/go-web-templates/api/internal/application/interfaces"
	"github.com/go-playground/validator/v10"
)

type PlaygroudJsonValidator struct {
	inner *validator.Validate
}

func NewPlaygroundJsonValidator() *PlaygroudJsonValidator {
	return &PlaygroudJsonValidator{
		inner: validator.New(),
	}
}

func (jv *PlaygroudJsonValidator) Validate(data interface{}) (bool, []interfaces.ValidateError) {
	errors := []interfaces.ValidateError {}

	err := jv.inner.Struct(data)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			errors = append(errors, interfaces.ValidateError{
				Tag: err.ActualTag(),
				Field: err.Field(),
				Value: err.Value(),
				Message: err.Error(),
			})
		}
	}

	return len(errors) == 0, errors
}
