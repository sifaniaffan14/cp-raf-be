package dto

import "cp-raf-be/internal/models"

type PageResponse struct {
	Title   string `json:"title"`
	Slug    string `json:"slug"`
}

func TransformPage(p models.Page) PageResponse {
	return PageResponse{
		Title: p.Title,
		Slug:  p.Slug,
	}
}