package main

import (
	"github.com/brunoleitem/go-img-api/api"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
	server := gin.Default()
	api.RegisterRoutes(server)
	server.Run(":3333")
}
