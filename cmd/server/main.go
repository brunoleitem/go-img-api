package main

import (
	"github.com/brunoleitem/go-wm/internal/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	routes.RegisterRoutes(server)
	server.Run(":3333")
}
