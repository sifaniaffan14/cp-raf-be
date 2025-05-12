package services

import (
	"cp-raf-be/database"
	"cp-raf-be/internal/models"
	"cp-raf-be/validators"
	"cp-raf-be/internal/dto"
	"fmt"
	"strconv"
)

type PageService struct{}

// CreatePage handles the logic for creating a new page
func (s *PageService) CreatePage(req validators.CreatePageRequest) (models.Page, error) {
	page := models.Page{
		Title:   req.Title,
		Content: req.Content,
		Slug:    req.Slug,
	}

	if err := database.DB.Create(&page).Error; err != nil {
		return models.Page{}, err
	}

	return page, nil
}

// GetPages retrieves all pages from the database
func (s *PageService) GetPages() ([]dto.PageResponse, error) {
	var pages []models.Page
	if err := database.DB.Find(&pages).Error; err != nil {
		return nil, err
	}

	// Mapping models.Page ke dto.PageResponse
	var result []dto.PageResponse
	for _, p := range pages {
		result = append(result, dto.TransformPage(p))  // Hanya memilih Title dan Slug
	}

	return result, nil
}

// UpdatePage updates an existing page based on the request
func (s *PageService) UpdatePage(req validators.UpdatePageRequest) (models.Page, error) {
	pageID, err := strconv.ParseInt(req.PageId, 10, 64)
	if err != nil {
		return models.Page{}, fmt.Errorf("invalid page ID format")
	}

	var page models.Page
	if err := database.DB.First(&page, pageID).Error; err != nil {
		return models.Page{}, fmt.Errorf("page not found")
	}

	if req.Title != "" {
		page.Title = req.Title
	}
	if req.Content != "" {
		page.Content = req.Content
	}
	if req.Slug != "" {
		page.Slug = req.Slug
	}

	if err := database.DB.Save(&page).Error; err != nil {
		return models.Page{}, err
	}

	return page, nil
}

// DeletePage deletes an existing page based on the request
func (s *PageService) DeletePage(req validators.DeletePageRequest) error {
	pageID, err := strconv.ParseInt(req.PageId, 10, 64)
	if err != nil {
		return fmt.Errorf("invalid page ID format")
	}

	var page models.Page
	if err := database.DB.First(&page, pageID).Error; err != nil {
		return fmt.Errorf("page not found")
	}

	if err := database.DB.Delete(&page).Error; err != nil {
		return err
	}

	return nil
}
