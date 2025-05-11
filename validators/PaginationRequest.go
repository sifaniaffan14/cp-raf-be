package validators

type PaginationRequest struct {
	PerPage string `json:"perPage" form:"perPage" validate:"required,numeric,excludesall=.,,"`
	Page    string `json:"page" form:"page" validate:"required,numeric,excludesall=.,,"`
}

func (r *PaginationRequest) Validate() error {
	return GlobalValidate(r)
}
