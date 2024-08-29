package main

import (
	"github.com/brunoleitem/go-img-api/api"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	api.RegisterRoutes(server)
	server.Run(":3333")
}
