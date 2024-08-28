package controllers

import (
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func Uploader(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Erro ao realizar upload", "error": err.Error()})
		return
	}

	ext := filepath.Ext(file.Filename)
	if ext != "jpg" && ext != "png" {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Erro ao realizar upload", "error": "Tipo de arquivo não suportado"})
		return
	}

	newFileName := uuid.New().String() + ext
	err = c.SaveUploadedFile(file, "tmp/"+newFileName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Erro ao realizar upload", "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Uploaded!"})
}
