package handlers

import (
	"context"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/brunoleitem/go-img-api/internal/img"
	"github.com/brunoleitem/go-img-api/internal/r2"
	"github.com/brunoleitem/go-img-api/internal/redis"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func Uploader(c *gin.Context) {
	r2, err := r2.NewR2Service()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Erro interno", "error": err.Error()})
		return
	}
	redisClient, err := redis.NewRedisService()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Erro interno", "error": err.Error()})
		return
	}
	texto := c.Query("text")
	if len(texto) > 110 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Erro", "error": "maximo de caracteres"})
		return
	}

	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Erro", "error": "erro ao recuperar o arquivo"})
		return
	}

	ext := strings.ToLower(filepath.Ext(file.Filename))
	if ext != ".jpg" && ext != ".jpeg" && ext != ".png" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Erro", "error": "tipo de arquivo inválido"})
		return
	}

	f, err := file.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Erro interno", "error": err.Error()})
		return
	}
	defer f.Close()

	loadedImg, err := img.LoadImage(f)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Erro interno", "error": err.Error()})
		return
	}

	processedImg, err := img.ProcessImage(loadedImg, ext, texto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Erro interno", "error": err.Error()})
		return
	}

	newFileName := uuid.New().String() + ext
	contentType := "image/" + strings.TrimPrefix(ext, ".")

	err = r2.UploadImage(context.TODO(), &newFileName, processedImg, contentType)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Erro interno", "error": err.Error()})
		return
	}
	userKey, err := redisClient.CreateImageKey(context.TODO(), newFileName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Erro interno", "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Sucesso", "filename": newFileName, "userKey": userKey})
}
