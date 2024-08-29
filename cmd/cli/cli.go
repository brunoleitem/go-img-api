package main

import (
	"context"

	"github.com/brunoleitem/go-img-api/internal/r2"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}

	r2service, err := r2.NewR2Service()
	if err != nil {
		panic(err)
	}
	r2service.ListBuckets(context.TODO())
	// i, err := img.LoadImage("image.png")
	// if err != nil {
	// 	panic(err)
	// }

	// newImg, err := img.ProcessImage(i)
	// if err != nil {
	// 	panic(err)
	// }

	// err = img.SaveImage("output.png", newImg)
	// if err != nil {
	// 	panic(err)
	// }
}
