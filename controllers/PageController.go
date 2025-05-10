package controllers

import (
	"cp-raf-be/database"
	"cp-raf-be/models"
	"cp-raf-be/utils"
	"cp-raf-be/validators"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreatePage(c *gin.Context) {
	var req validators.CreatePageRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.SendError(c, "Invalid JSON", "VALIDATION_ERROR", http.StatusBadRequest)
		return
	}

	if err := req.Validate(); err != nil {
		utils.SendError(c, err.Error(), "VALIDATION_ERROR", http.StatusBadRequest)
		return
	}

	page := models.Page{
		Title:   req.Title,
		Content: req.Content,
		Slug:    req.Slug,
	}

	if err := database.DB.Create(&page).Error; err != nil {
		utils.SendError(c, err.Error(), "DB_ERROR", http.StatusInternalServerError)
		return
	}

	utils.SendResponse(c, page, "Page created successfully", http.StatusCreated)
}

func GetPages(c *gin.Context) {
	var pages []models.Page
	if err := database.DB.Find(&pages).Error; err != nil {
		utils.SendError(c, "Failed to fetch pages", "500", http.StatusInternalServerError)
		return
	}

	// Jika pages kosong, kirimkan response error
	if len(pages) == 0 {
		utils.SendError(c, "No pages found", "404", http.StatusNotFound)
		return
	}

	// Kirimkan response dengan data yang ditemukan
	utils.SendResponse(c, pages, "Pages retrieved successfully", http.StatusOK)
}
