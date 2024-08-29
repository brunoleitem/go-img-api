package main

import "github.com/brunoleitem/go-img-api/internal/img"

func main() {
	i, err := img.LoadImage("image.png")
	if err != nil {
		panic(err)
	}

	newImg, err := img.ProcessImage(i)
	if err != nil {
		panic(err)
	}

	err = img.SaveImage("output.png", newImg)
	if err != nil {
		panic(err)
	}
}
