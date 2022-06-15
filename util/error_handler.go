package util

import (
	"fmt"
	"net/http"
	"reflect"
	"strings"

	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/krobus00/technical-test-rest-api/model"
	"github.com/labstack/echo/v4"
)

func RegisterTagNameWithLabel(validate *validator.Validate) {
	validate.RegisterTagNameFunc(func(field reflect.StructField) string {
		name := strings.SplitN(field.Tag.Get("label"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})
}

func RegisterTagNameWithJson(validate *validator.Validate) {
	validate.RegisterTagNameFunc(func(field reflect.StructField) string {
		name := strings.SplitN(field.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})
}

func BuildValidationErrors(err error, trans ut.Translator) []model.ValidationError {
	errors := []model.ValidationError{}
	for _, err := range err.(validator.ValidationErrors) {
		errors = append(errors, model.ValidationError{
			Field:   strings.Join(strings.Split(err.Namespace(), ".")[1:], "."),
			Message: err.Translate(trans),
		})
	}
	return errors
}

func CustomHTTPErrorHandler(err error, c echo.Context) {
	var (
		response model.RequestError
	)

	// default
	statusCode := http.StatusInternalServerError
	message := "Internal Server Error"

	if e, ok := err.(*model.HttpCustomError); ok {
		statusCode = e.StatusCode
		response.Error = model.HttpCustomError{
			StatusCode: e.StatusCode,
			Message:    e.Message,
		}
	} else if e, ok := err.(*echo.HTTPError); ok {
		statusCode = e.Code
		if m, ok := e.Message.([]model.ValidationError); ok {
			response.Error = model.HttpCustomErrors{
				StatusCode: statusCode,
				Message:    "validation errors",
				Errors:     m,
			}
		} else {
			response.Error = model.HttpCustomError{
				StatusCode: statusCode,
				Message:    fmt.Sprintf("%s", e.Message),
			}
		}
	} else {
		response.Error = model.HttpCustomError{
			StatusCode: http.StatusInternalServerError,
			Message:    message,
		}
	}

	c.JSON(statusCode, response)
}
