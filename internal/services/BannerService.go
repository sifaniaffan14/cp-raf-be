package services

import (
	"cp-raf-be/database"
	"cp-raf-be/internal/models"
	"cp-raf-be/validators"
	"cp-raf-be/internal/dto"
	"fmt"
	"strconv"
)

type BannerService struct{}

// CreateBanner handles the logic for creating a new page
func (s *BannerService) CreateBanner(req validators.CreateBannerRequest) (models.Banner, error) {
	banner := models.Banner{
		ImageUrl:   req.ImageUrl,
		Tittle: req.Tittle,
		Description:    req.Description,
	}

	if err := database.DB.Create(&banner).Error; err != nil {
		return models.Banner{}, err
	}

	return banner, nil
}

// GetBanners retrieves all pages from the database
func (s *BannerService) GetBanners() ([]dto.BannerResponse, error) {
	var banners []models.Banner
	if err := database.DB.Find(&banners).Error; err != nil {
		return nil, err
	}

	// Mapping models.Banner ke dto.BannerResponse
	var result []dto.BannerResponse
	for _, p := range banners {
		result = append(result, dto.TransformBanner(p)) 
	}

	return result, nil
}

// UpdateBanner updates an existing page based on the request
func (s *BannerService) UpdateBanner(req validators.UpdateBannerRequest) (models.Banner, error) {
	bannerId, err := strconv.ParseInt(req.BannerId, 10, 64)
	if err != nil {
		return models.Banner{}, fmt.Errorf("invalid banner ID format")
	}

	var banner models.Banner
	if err := database.DB.First(&banner, bannerId).Error; err != nil {
		return models.Banner{}, fmt.Errorf("page not found")
	}

	if req.ImageUrl != "" {
		banner.ImageUrl = req.ImageUrl
	}
	if req.Tittle != "" {
		banner.Tittle = req.Tittle
	}
	if req.Description != "" {
		banner.Description = req.Description
	}

	if err := database.DB.Save(&banner).Error; err != nil {
		return models.Banner{}, err
	}

	return banner, nil
}

// DeleteBanner deletes an existing page based on the request
func (s *BannerService) DeleteBanner(req validators.DeleteBannerRequest) error {
	bannerId, err := strconv.ParseInt(req.BannerId, 10, 64)
	if err != nil {
		return fmt.Errorf("invalid banner ID format")
	}

	var banner models.Banner
	if err := database.DB.First(&banner, bannerId).Error; err != nil {
		return fmt.Errorf("banner not found")
	}

	if err := database.DB.Delete(&banner).Error; err != nil {
		return err
	}

	return nil
}
