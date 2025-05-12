package validators

import (
	"github.com/go-playground/validator/v10"
)

// PaginationRequest digunakan untuk menerima parameter pagination dari query atau body
type PaginationRequest struct {
	Page    string `json:"page" form:"page" validate:"required,number"`
	PerPage string `json:"perPage" form:"perPage" validate:"required,number"`
}

// Validate memvalidasi data pagination
func (p *PaginationRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(p)
}