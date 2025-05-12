package controllers

import (
	"cp-raf-be/internal/services"
	"cp-raf-be/utils"
	"cp-raf-be/validators"
	"github.com/gin-gonic/gin"
	"net/http"
	"cp-raf-be/helpers"
)

var bannerService = services.BannerService{}
	// Ambil data dari service
func CreateBanner(c *gin.Context) {
	tittle := c.PostForm("tittle")
	description := c.PostForm("description")

	file, err := c.FormFile("image")
	if err != nil {
		utils.SendError(c, "Image is required", "VALIDATION_ERROR", http.StatusBadRequest)
		return
	}

	// Upload ke S3/Minio
	imageUrl, err := helpers.UploadToS3(file)
	if err != nil {
		utils.SendError(c, "Failed to upload image: "+err.Error(), "UPLOAD_ERROR", http.StatusInternalServerError)
		return
	}

	req := validators.CreateBannerRequest{
		ImageUrl:   imageUrl,
		Tittle:     tittle,
		Description: description,
	}

	if err := req.Validate(); err != nil {
		utils.SendError(c, err.Error(), "VALIDATION_ERROR", http.StatusBadRequest)
		return
	}

	banner, err := bannerService.CreateBanner(req)
	if err != nil {
		utils.SendError(c, err.Error(), "DB_ERROR", http.StatusInternalServerError)
		return
	}

	responseData := gin.H{
		"id":          banner.Id,
		"imageUrl":    banner.ImageUrl,
		"tittle":      banner.Tittle,
		"description": banner.Description,
	}

	utils.SendResponse(c, responseData, "Banner created successfully", http.StatusCreated)
}