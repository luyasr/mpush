package validate

import (
	"errors"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhTrans "github.com/go-playground/validator/v10/translations/zh"
	"reflect"
)

func Struct(obj any) error {
	validate := validator.New()

	// english translator
	enTranslator := en.New()
	// chinese translator
	zhTranslator := zh.New()
	uni := ut.New(enTranslator, zhTranslator)
	// get the required language
	trans, _ := uni.GetTranslator("zhTranslator")
	// register a function to get the custom label in the struct tag as the field name
	validate.RegisterTagNameFunc(func(field reflect.StructField) string {
		name := field.Tag.Get("label")
		return name
	})

	// validator registration translator
	err := zhTrans.RegisterDefaultTranslations(validate, trans)
	if err != nil {
		return err
	}
	err = validate.Struct(obj)
	if err != nil {
		var invalidValidationError *validator.InvalidValidationError
		if errors.As(err, &invalidValidationError) {
			return errors.New(err.Error())
		}
		for _, e := range err.(validator.ValidationErrors) {
			return errors.New(e.Translate(trans))
		}
	}

	return nil
}
