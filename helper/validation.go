package helper

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"
)

func FormatValidationError(err error) []string {
	var errorsList []string

	// Check if the error is of type ValidationErrors
	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		for _, e := range validationErrors {
			errorsList = append(errorsList, e.Error())
		}
		return errorsList
	}

	// Check if the error is a JSON UnmarshalTypeError
	var unmarshalTypeError *json.UnmarshalTypeError
	if errors.As(err, &unmarshalTypeError) {
		errorsList = append(errorsList, fmt.Sprintf("Field '%s' has an invalid type", unmarshalTypeError.Field))
		return errorsList
	}

	// Check if the error is a JSON SyntaxError
	var syntaxError *json.SyntaxError
	if errors.As(err, &syntaxError) {
		errorsList = append(errorsList, "Invalid JSON syntax")
		return errorsList
	}

	// Handle general errors
	errorsList = append(errorsList, err.Error())
	return errorsList
}

