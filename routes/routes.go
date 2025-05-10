package routes

import (
	"cp-raf-be/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	router.GET("/admin/pages", controllers.GetPages)
	router.POST("/admin/pages", controllers.CreatePage)
}
