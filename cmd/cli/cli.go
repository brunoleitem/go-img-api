package main

import (
	"context"
	"log"

	"github.com/brunoleitem/go-img-api/internal/r2"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
	rt, err := r2.NewR2Service()
	if err != nil {
		panic(err)
	}
	lala := "6f1a4b2a-1105-41bb-bd13-a26bc612e3d0.jpg"
	var s *string = &lala

	err = rt.DeleteImage(context.TODO(), s)
	if err != nil {
		log.Fatalf("ERRO DELETANDO %v", err)
	}
}
