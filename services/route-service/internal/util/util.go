package util

import (
	"github.com/go-playground/validator/v10"
	"reflect"
)

type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

// ErrorResponse represents the structure of error responses.
type ErrorResponse struct {
	Code    int               `json:"code"`
	Message string            `json:"message"`
	Errors  []ValidationError `json:"errors,omitempty"`
}

// FormatValidationError formats validation errors from Gin into a slice of ValidationError.
// It uses the json tags of the dto to return field names as they are in the json request.
func FormatValidationError(err error, dtoType interface{}) []ValidationError {
	var errors []ValidationError
	valErrs := err.(validator.ValidationErrors)
	dtoVal := reflect.TypeOf(dtoType)

	for _, fe := range valErrs {
		field, found := dtoVal.FieldByName(fe.StructField())
		jsonTag := field.Tag.Get("json")
		fieldName := jsonTag
		if found && jsonTag == "" {
			fieldName = field.Name
		}

		errors = append(errors, ValidationError{
			Field:   fieldName,
			Message: fe.ActualTag(),
		})
	}
	return errors
}
