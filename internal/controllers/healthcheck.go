package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleHealthCheck(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "OK"})
}
