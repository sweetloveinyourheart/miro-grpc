package utils

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

// This is the validator instance
// for more information see: https://github.com/go-playground/validator
var validate = validator.New()

type ErrorResponse struct {
	Error       bool
	FailedField string
	Tag         string
	Param       string
	Value       any
}

func Validate(data any) []ErrorResponse {
	validationErrors := []ErrorResponse{}

	errs := validate.Struct(data)
	if errs != nil {
		for _, err := range errs.(validator.ValidationErrors) {
			var elem ErrorResponse

			elem.FailedField = err.Field() // Export struct field name
			elem.Tag = err.Tag()           // Export struct tag
			elem.Value = err.Value()       // Export field value
			elem.Param = err.Param()       // Export param
			elem.Error = true

			validationErrors = append(validationErrors, elem)
		}
	}

	return validationErrors
}

func CreateValidationMessage(errs []ErrorResponse) string {
	translatedErrors := translateErrors(errs)
	message := strings.Join(translatedErrors, ", ")
	return message
}

func translateErrors(errs []ErrorResponse) []string {
	var errorMessages []string
	for _, err := range errs {
		switch err.Tag {
		case "required":
			errorMessages = append(errorMessages, fmt.Sprintf("[%s] is required", err.FailedField))
		case "min":
			errorMessages = append(errorMessages, fmt.Sprintf("[%s] must be at least %s characters long", err.FailedField, err.Param))
		case "max":
			errorMessages = append(errorMessages, fmt.Sprintf("[%s] must be at most %s characters long", err.FailedField, err.Param))
		case "email":
			errorMessages = append(errorMessages, fmt.Sprintf("[%s] must be a valid email address", err.FailedField))
		case "gte":
			errorMessages = append(errorMessages, fmt.Sprintf("[%s] must be greater than or equal to %s", err.FailedField, err.Param))
		case "lte":
			errorMessages = append(errorMessages, fmt.Sprintf("[%s] must be less than or equal to %s", err.FailedField, err.Param))
		default:
			errorMessages = append(errorMessages, fmt.Sprintf("[%s] is not valid", err.FailedField))
		}
	}
	return errorMessages
}
