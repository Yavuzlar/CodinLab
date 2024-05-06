package validator_service

import (
	"log"
	"reflect"
	"strings"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

type ValidatorService struct {
	validate *validator.Validate
}

type keyErr struct {
	Key string `json:"field"`
	Err string `json:"error"`
}

// func GetTranslator(lang string) *ut.Translator {
// 	switch lang {
// 	case "tr":
// 		return transTr
// 	case "en":
// 		return transEn
// 	default:
// 		return transEn
// 	}
// }

// NewValidator func for create a new validator for model fields.

func NewValidatorService() *ValidatorService {

	enLang := en.New()
	uni := ut.New(enLang, enLang)
	transen, _ := uni.GetTranslator("en")
	validate := validator.New()
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		// skip if tag key says it should be ignored
		if name == "-" {
			return ""
		}
		return name
	})
	if err := en_translations.RegisterDefaultTranslations(validate, transen); err != nil {
		log.Fatalf("error on register default en translations")
	}
	return &ValidatorService{
		validate: validate,
	}
}

func (vs *ValidatorService) ValidateStruct(s any) error {
	return vs.validate.Struct(s)
}

// ValidatorErrors func for show validation errors for each invalid fields.
func ValidatorErrors(err error) (result []keyErr) {
	for _, err := range err.(validator.ValidationErrors) {
		result = append(result, keyErr{
			Key: err.Field(),
			Err: err.Error(),
		})
	}
	return result
}
