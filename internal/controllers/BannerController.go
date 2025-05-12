package controllers

import (
	"cp-raf-be/internal/services"
	"cp-raf-be/utils"
	"cp-raf-be/validators"
	"github.com/gin-gonic/gin"
	"net/http"
)
	// Ambil data dari service
func CreateBanner(c *gin.Context) {
	var req validators.CreatePageRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.SendError(c, "Invalid JSON", "VALIDATION_ERROR", http.StatusBadRequest)
		return
	}

	if err := req.Validate(); err != nil {
		utils.SendError(c, err.Error(), "VALIDATION_ERROR", http.StatusBadRequest)
		return
	}

	// Memanggil service untuk membuat halaman
	pageService := services.PageService{}
	page, err := pageService.CreatePage(req)
	if err != nil {
		utils.SendError(c, err.Error(), "DB_ERROR", http.StatusInternalServerError)
		return
	}

	// Hanya mengembalikan data yang dibutuhkan
	responseData := gin.H{
		"id":      page.ID,
		"title":   page.Title,
		"content": page.Content,
		"slug":    page.Slug,
	}

	utils.SendResponse(c, responseData, "Page created successfully", http.StatusCreated)
}