package helper

import (
	"github.com/go-playground/validator/v10"
)

type Response struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

func APIResponse(status string, data interface{}) Response {
	jsonResponse := Response{
		Status: status,
		Data:   data,
	}

	return jsonResponse
}

func FormatValidationError(err error) []string {
	var errors []string

	// loop errors
	for _, e := range err.(validator.ValidationErrors) {
		errors = append(errors, e.Error())
	}

	return errors
}
