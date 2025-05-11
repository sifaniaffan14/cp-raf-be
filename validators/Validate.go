package validators

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)


var Validate *validator.Validate

func InitValidator() {
	Validate = validator.New()
}

func GlobalValidate(i interface{}) error {
	err := Validate.Struct(i)
	if err == nil {
		return nil
	}

	var errorMessages []string
	for _, err := range err.(validator.ValidationErrors) {
		errorMessages = append(errorMessages, fmt.Sprintf("Field '%s' is %s", err.Field(), err.Tag()))
	}

	return fmt.Errorf("validation failed: %s", strings.Join(errorMessages, ", "))
}