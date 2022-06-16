package infrastructure

import (
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/id"
	"github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	idTranslations "github.com/go-playground/validator/v10/translations/id"
	"github.com/krobus00/technical-test-rest-api/util"
	"github.com/labstack/echo/v4"
)

type (
	ValidationUtil struct {
		validator *validator.Validate
	}
)

func NewTranslator() *ut.UniversalTranslator {
	english := en.New()
	indo := id.New()
	uni := ut.New(indo, english)
	return uni
}

func NewValidator(trans *ut.UniversalTranslator) echo.Validator {
	validate := validator.New()
	util.RegisterTagNameWithJson(validate)

	id, _ := trans.GetTranslator("id")
	en, _ := trans.GetTranslator("en")
	_ = enTranslations.RegisterDefaultTranslations(validate, en)
	_ = idTranslations.RegisterDefaultTranslations(validate, id)

	return &ValidationUtil{validator: validate}
}

func (v *ValidationUtil) Validate(i interface{}) error {
	return v.validator.Struct(i)
}
