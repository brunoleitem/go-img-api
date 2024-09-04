package redis

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

type RedisService struct {
	client *redis.Client
}

func NewRedisService() (*RedisService, error) {
	url := os.Getenv("REDIS_URL")
	user := os.Getenv("REDIS_USER")
	password := os.Getenv("REDIS_PASSWORD")
	db, err := strconv.Atoi(os.Getenv("REDIS_DB"))
	if err != nil {
		return nil, errors.New("erro ao conectar com redis DB")
	}

	client := redis.NewClient(&redis.Options{
		Addr:     url,
		Password: password,
		Username: user,
		DB:       db,
		OnConnect: func(ctx context.Context, cn *redis.Conn) error {
			fmt.Println("Redis connected")
			return nil
		},
	})

	return &RedisService{
		client: client,
	}, nil
}

func (r *RedisService) CreateImageKey(ctx context.Context, imgKey string) (string, error) {
	userPassKey := uuid.New().String()
	err := r.client.Set(ctx, imgKey, userPassKey, 168*time.Hour)
	if err.Err() != nil {
		return "", err.Err()
	}
	return userPassKey, nil
}
