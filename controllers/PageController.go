package controllers

import (
	"cp-raf-be/services"
	"cp-raf-be/utils"
	"cp-raf-be/validators"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
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

func GetPages(c *gin.Context) {
	// Ambil data dari service
	pages, err := pageService.GetPages()
	if err != nil {
		utils.SendError(c, err.Error(), "DB_ERROR", http.StatusInternalServerError)
		return
	}

	if len(pages) == 0 {
		utils.SendError(c, "No pages found", "NOT_FOUND", http.StatusNotFound)
		return
	}

	// Gunakan struct dari validators untuk pagination request
	var pagination validators.PaginationRequest
	_ = c.ShouldBindJSON(&pagination)

	// Prioritaskan dari payload JSON, fallback ke query string
	perPage := c.DefaultQuery("perPage", "15")
	if pagination.PerPage != "" {
		perPage = pagination.PerPage
	}
	page := c.DefaultQuery("page", "1")
	if pagination.Page != "" {
		page = pagination.Page
	}

	// Konversi ke integer
	perPageInt, err := strconv.Atoi(perPage)
	if err != nil || perPageInt <= 0 {
		perPageInt = 15
	}
	pageInt, err := strconv.Atoi(page)
	if err != nil || pageInt <= 0 {
		pageInt = 1
	}

	// Persiapkan hasil dengan merubah data yang akan dipaginate
	var result []gin.H
	for _, p := range pages {
		result = append(result, gin.H{
			"id":      p.ID,
			"title":   p.Title,
			"content": p.Content,
			"slug":    p.Slug,
		})
	}

	// Paginate data dan mengembalikan hasil
	paginatedData, meta := utils.Paginate(c, result, pageInt, perPageInt)

	// Kirim response hasil paginasi dengan meta data
	utils.SendResponse(c, gin.H{
		"pages": paginatedData,
		"meta": meta,
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
