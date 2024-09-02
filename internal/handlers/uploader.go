package handlers

import (
	"context"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/brunoleitem/go-img-api/internal/img"
	"github.com/brunoleitem/go-img-api/internal/r2"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func Uploader(c *gin.Context) {
	r2, err := r2.NewR2Service()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to initialize R2 service", "error": err.Error()})
		return
	}

	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Failed to get file from form", "error": err.Error()})
		return
	}

	ext := strings.ToLower(filepath.Ext(file.Filename))
	if ext != ".jpg" && ext != ".jpeg" && ext != ".png" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Unsupported file type"})
		return
	}

	f, err := file.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to open file", "error": err.Error()})
		return
	}
	defer f.Close()

	loadedImg, err := img.LoadImage(f)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to decode image", "error": err.Error()})
		return
	}

	processedImg, err := img.ProcessImage(loadedImg, ext)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to process image", "error": err.Error()})
		return
	}

	newFileName := uuid.New().String() + ext
	contentType := "image/" + strings.TrimPrefix(ext, ".")

	err = r2.UploadImage(context.TODO(), &newFileName, processedImg, contentType)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to upload image", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Image uploaded successfully", "filename": newFileName})
}
