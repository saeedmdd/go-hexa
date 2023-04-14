package http

import (
	"github.com/go-playground/validator/v10"
	"github.com/saeedmdd/go-hexa/internal/core/models"
)

type ErrorResponse struct {
	FailedField string `json:"field"`
	Tag         string `json:"tag"`
	Value       string `json:"value"`
}

var validate = validator.New()

func ValidateTask(task models.TaskModel) (errors []*ErrorResponse) {
	err := validate.Struct(task)
	if err != nil {
		for _, fieldError := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.FailedField = fieldError.StructNamespace()
			element.Tag = fieldError.Tag()
			element.Value = fieldError.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}
