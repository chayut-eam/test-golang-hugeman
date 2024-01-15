package validation

import (
	"fmt"
	"reflect"
	"strings"

	e "github.com/chayut-eam/test-golang-hugeman/error"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	v "github.com/go-playground/validator/v10"
)

var (
	validator *v.Validate
	trans     ut.Translator
)

func Init() {
	en := en.New()
	uni := ut.New(en, en)
	trans, _ = uni.GetTranslator("en")

	validator = v.New()

	// register custom validator
	validator.RegisterValidation("notEmpty", notEmptyValidator)

	// register tag name
	registerTagName("json")

	// register message translate
	registerTranslation("required", "is required")
	registerTranslation("max", "more than %v characters")
	registerTranslation("oneof", "must be one of %v")
}

func Validate(s interface{}) error {
	if err := validator.Struct(s); err != nil {
		if errs, ok := err.(v.ValidationErrors); ok {
			messages := map[string]string{}
			for _, e := range errs {
				messages[e.Field()] = e.Translate(trans)
			}
			return e.NewFieldValidationError(messages)
		}
		return err
	}
	return nil
}

func registerTagName(tags ...string) {
	for _, tag := range tags {
		validator.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get(tag), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})
	}
}

func registerTranslation(tag string, errmsg string) {
	registerFn := func(ut ut.Translator) error {
		return ut.Add(tag, errmsg, false)
	}

	transFn := func(ut ut.Translator, fe v.FieldError) string {
		param := fe.Param()
		tag := fe.Tag()

		t, err := ut.T(tag, fe.Field(), param)
		if err != nil {
			return fe.(error).Error()
		}

		if param == "" {
			return t
		}

		return fmt.Sprintf(t, param)
	}

	_ = validator.RegisterTranslation(tag, trans, registerFn, transFn)
}

func notEmptyValidator(fl v.FieldLevel) bool {
	value := fl.Field().String()
	return len(value) > 0
}
