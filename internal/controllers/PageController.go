package controllers

import (
	"cp-raf-be/internal/services"
	"cp-raf-be/utils"
	"cp-raf-be/validators"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

var pageService = services.PageService{}

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

func dd(v interface{}) {
	fmt.Printf("%+v\n", v)
	os.Exit(1)
}

func GetPages(c *gin.Context) {
	pages, err := pageService.GetPages()
	if err != nil {
		utils.SendError(c, err.Error(), "DB_ERROR", http.StatusInternalServerError)
		return
	}

	if len(pages) == 0 {
		utils.SendError(c, "No pages found", "NOT_FOUND", http.StatusNotFound)
		return
	}

	page, perPage := utils.GetPaginationParams(c)
	paginatedData, meta := utils.Paginate(c, pages, page, perPage)

	utils.SendResponse(c, gin.H{
		"pages": paginatedData,
		"meta":  meta,
	}, "Pages retrieved successfully", http.StatusOK)
}
func UpdatePage(c *gin.Context) {
	var req validators.UpdatePageRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.SendError(c, "Invalid JSON", "VALIDATION_ERROR", http.StatusBadRequest)
		return
	}

	if err := req.Validate(); err != nil {
		utils.SendError(c, err.Error(), "VALIDATION_ERROR", http.StatusBadRequest)
		return
	}

	page, err := pageService.UpdatePage(req)
	if err != nil {
		utils.SendError(c, err.Error(), "DB_ERROR", http.StatusInternalServerError)
		return
	}

	utils.SendResponse(c, gin.H{"id": page.ID}, "Page updated successfully", http.StatusOK)
}

func DeletePage(c *gin.Context) {
	var req validators.DeletePageRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.SendError(c, "Invalid JSON", "VALIDATION_ERROR", http.StatusBadRequest)
		return
	}

	if err := req.Validate(); err != nil {
		utils.SendError(c, err.Error(), "VALIDATION_ERROR", http.StatusBadRequest)
		return
	}

	err := pageService.DeletePage(req)
	if err != nil {
		utils.SendError(c, err.Error(), "DB_ERROR", http.StatusInternalServerError)
		return
	}

	utils.SendResponse(c, gin.H{"id": req.PageId}, "Page deleted successfully", http.StatusOK)
}
