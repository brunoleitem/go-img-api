package api

import (
	"github.com/brunoleitem/go-img-api/internal/handlers"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/api/healthcheck", handlers.HandleHealthCheck)
	server.POST("/api/upload", handlers.Uploader)
}
