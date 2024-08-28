package routes

import (
	"github.com/brunoleitem/go-wm/internal/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/api/healthcheck", controllers.HandleHealthCheck)
	server.POST("/api/upload", controllers.Uploader)
}
