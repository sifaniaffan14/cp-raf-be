package dto

import "cp-raf-be/internal/models"

type BannerResponse struct {
	ImageUrl    string `json:"image_url"`
	Tittle      string `json:"tittle"`
	Description string `json:"description"`
}

func TransformBanner(p models.Banner) BannerResponse {
	return BannerResponse{
		ImageUrl:    p.ImageUrl,
		Tittle:      p.Tittle,
		Description: p.Description,
	}
}
