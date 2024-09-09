package handlers

import (
	"context"
	"net/http"

	"github.com/brunoleitem/go-img-api/internal/r2"
	"github.com/brunoleitem/go-img-api/internal/redis"
	"github.com/gin-gonic/gin"
)

func Delete(c *gin.Context) {
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

	userKey := c.Param("id")

	imgId, err := redisClient.GetKeyValue(context.TODO(), &userKey)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Erro", "error": "Chave de usuario inválida"})
		return
	}

	err = r2.DeleteImage(context.TODO(), &imgId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Erro", "error": "Erro ao excluir imagem"})
		return
	}

	err = redisClient.DeleteKey(context.TODO(), &userKey)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Erro", "error": "Erro ao excluir imagem"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Sucesso"})
}
