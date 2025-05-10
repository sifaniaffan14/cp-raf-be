package validators

import (
	"fmt"
	"strings"

	"cp-raf-be/models"
	"github.com/go-playground/validator/v10"
)

var Validate *validator.Validate

func InitValidator() {
	Validate = validator.New()
}

type CreatePageRequest struct {
	Title   string `json:"title" validate:"required,min=1"`
	Content string `json:"content" validate:"required,min=1"`
	Slug    string `json:"slug" validate:"required,min=1"`
}

func (r *CreatePageRequest) Validate() error {
	err := Validate.Struct(r)
	if err == nil {
		return nil
	}

	var errorMessages []string
	for _, err := range err.(validator.ValidationErrors) {
		errorMessages = append(errorMessages, fmt.Sprintf("Field '%s' is %s", err.Field(), err.Tag()))
	}

	return fmt.Errorf("validation failed: %s", strings.Join(errorMessages, ", "))
}

func MapPageToRequest(page models.Page) CreatePageRequest {
	return CreatePageRequest{
		Title:   page.Title,
		Content: page.Content,
		Slug:    page.Slug,
	}
}