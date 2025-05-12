package validators

import (
	"cp-raf-be/internal/models"
	"errors"
	"mime/multipart"
	"path/filepath"
	"strings"
)

type CreateBannerRequest struct {
	Image       *multipart.FileHeader `form:"image"`
	ImageUrl    string                `json:"-"` // optional: tidak dikirim dari client
	Tittle      string                `form:"tittle" validate:"required,min=1"`
	Description string                `form:"description"`
}

// Validasi manual untuk file image
func (r *CreateBannerRequest) Validate() error {
	if r.Image == nil {
		return errors.New("image is required")
	}

	ext := strings.ToLower(filepath.Ext(r.Image.Filename))
	if ext != ".jpg" && ext != ".jpeg" && ext != ".png" {
		return errors.New("only JPG, JPEG, and PNG formats are allowed")
	}

	const maxSize = 5 << 20 // 5MB
	if r.Image.Size > maxSize {
		return errors.New("image size must be less than 5MB")
	}

	return GlobalValidate(r)
}

func MapBannerToCreateRequest(banner models.Banner) CreateBannerRequest {
	return CreateBannerRequest{
		ImageUrl:    banner.ImageUrl,
		Tittle: banner.Tittle,
		Description:    banner.Description,
	}
}

type UpdateBannerRequest struct {
	BannerId  string `json:"bannerId" validate:"required"`
	ImageUrl   string `json:"image_url"`
	Tittle string `json:"tittle"`
	Description    string `json:"description"`
}

func (r *UpdateBannerRequest) Validate() error {
	return GlobalValidate(r)
}

func MapBannerToUpdateRequest(banner models.Banner) UpdateBannerRequest {
    return UpdateBannerRequest{
        BannerId:  banner.Id,
        ImageUrl:   banner.ImageUrl,
        Tittle:     banner.Tittle,
        Description: banner.Description,
    }
}

type DeleteBannerRequest struct {
	BannerId string `json:"bannerId" validate:"required"`
}
func (r *DeleteBannerRequest) Validate() error {
	return GlobalValidate(r)
}

func MapBannerToDeleteRequest(banner models.Banner) DeleteBannerRequest {
    return DeleteBannerRequest{
        BannerId: banner.Id,
    }
}
