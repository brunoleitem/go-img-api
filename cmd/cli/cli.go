package main

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"math"
	"os"

	"github.com/golang/freetype"
)

func main() {
	img, err := loadImage("image.png")
	if err != nil {
		panic(err)
	}

	newImg, err := processImage(img)
	if err != nil {
		panic(err)
	}

	err = saveImage("output.png", newImg)
	if err != nil {
		panic(err)
	}
}

func loadImage(filePath string) (image.Image, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	img, _, err := image.Decode(f)
	if err != nil {
		return nil, err
	}

	return img, nil
}

func saveImage(filePath string, img *image.RGBA) error {
	out, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer out.Close()

	return png.Encode(out, img)
}

func processImage(img image.Image) (*image.RGBA, error) {
	newImg := drawRect(img)
	err := drawText(newImg)
	if err != nil {
		return nil, err
	}
	return newImg, nil
}

func drawText(newImg *image.RGBA) error {
	fontBytes, err := os.ReadFile("assets/arial.ttf")
	if err != nil {
		return err
	}
	c := freetype.NewContext()
	font, err := freetype.ParseFont(fontBytes)
	if err != nil {
		return err
	}
	c.SetDPI(72)
	c.SetFont(font)
	c.SetFontSize(24)
	c.SetClip(newImg.Bounds())
	c.SetDst(newImg)
	c.SetSrc(image.NewUniform(color.White))

	pt := freetype.Pt(25, newImg.Rect.Dy()-35)
	_, err = c.DrawString("Hello, World!", pt)
	if err != nil {
		return err
	}
	return nil
}

func drawRect(img image.Image) *image.RGBA {
	width := img.Bounds().Dx()
	height := img.Bounds().Dy()

	newImg := image.NewRGBA(image.Rect(0, 0, width, height))
	draw.Draw(newImg, newImg.Bounds(), img, image.Point{0, 0}, draw.Over)

	rectW := int(math.Round(float64(width) * 0.3))
	rectH := 60
	padding := 15
	// rectPos := image.Point{0 + padding, (height - rectH) - padding}
	rectPos := image.Rect(padding, height-rectH-padding, padding+rectW, height-padding)

	// rectImg := image.NewRGBA(image.Rect(0, 0, rectW, rectH))
	fillColor := color.RGBA{0, 0, 0, 64}
	draw.Draw(newImg, rectPos, &image.Uniform{fillColor}, image.Point{}, draw.Over)

	// draw.Draw(rectImg, rectImg.Bounds(), &image.Uniform{fillColor}, image.Point{0, 0}, draw.Src)

	// draw.Draw(newImg, image.Rect(rectPos.X, rectPos.Y, rectPos.X+rectW, rectPos.Y+rectH), rectImg, image.Point{0, 0}, draw.Over)
	return newImg
}
