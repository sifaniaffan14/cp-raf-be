package validators

import (
	"cp-raf-be/models"
	"fmt"
)

// Struct untuk create page
type CreatePageRequest struct {
	Title   string `json:"title" validate:"required,min=1"`
	Content string `json:"content" validate:"required,min=1"`
	Slug    string `json:"slug" validate:"required,min=1"`
}

func (r *CreatePageRequest) Validate() error {
	return GlobalValidate(r)
}

func MapPageToRequest(page models.Page) CreatePageRequest {
	return CreatePageRequest{
		Title:   page.Title,
		Content: page.Content,
		Slug:    page.Slug,
	}
}

// Struct untuk update page
type UpdatePageRequest struct {
	PageId  string `json:"pageId" validate:"required"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Slug    string `json:"slug"`
}

func (r *UpdatePageRequest) Validate() error {
	return GlobalValidate(r)
}

func UpdatePageToRequest(page models.Page) UpdatePageRequest {
	return UpdatePageRequest{
		PageId:  fmt.Sprintf("%d", page.ID),
		Title:   page.Title,
		Content: page.Content,
		Slug:    page.Slug,
	}
}

type DeletePageRequest struct {
	PageId string `json:"pageId" validate:"required"`
}
func (r *DeletePageRequest) Validate() error {
	return GlobalValidate(r)
}
/*************  ✨ Windsurf Command ⭐  *************/
// MapPageToDeleteRequest maps a models.Page to a DeletePageRequest.
// It only includes the ID of the page.
/*******  37db1f59-9b96-43f1-91e7-8b5d6f468c00  *******/
func MapPageToDeleteRequest(page models.Page) DeletePageRequest {
	return DeletePageRequest{
		PageId: fmt.Sprintf("%d", page.ID),
	}
}
