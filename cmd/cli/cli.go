package main

import (
	"os"
	"path/filepath"

	"github.com/brunoleitem/go-img-api/internal/img"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
	f, err := os.Open("image.png")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	i, err := img.LoadImage(f)
	if err != nil {
		panic(err)
	}
	newImg, err := img.ProcessImage(i)
	if err != nil {
		panic(err)
	}
	ext := filepath.Ext(f.Name())
	err = img.SaveImage(newImg, "output"+ext)
	if err != nil {
		panic(err)
	}
}
