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
	f, err := os.Open("image.png")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	img, _, err := image.Decode(f)
	if err != nil {
		panic(err)
	}

	newImg := drawRect(img)
	drawText(newImg)

	// Save the new image to a file
	out, err := os.Create("output.png")
	if err != nil {
		panic(err)
	}
	defer out.Close()

	err = png.Encode(out, newImg)
	if err != nil {
		panic(err)
	}
}

func drawText(newImg *image.RGBA) {
	fontBytes, err := os.ReadFile("assets/arial.ttf")
	if err != nil {
		panic(err)
	}
	c := freetype.NewContext()
	font, err := freetype.ParseFont(fontBytes)
	if err != nil {
		panic(err)
	}
	c.SetDPI(72)
	c.SetFont(font)
	c.SetFontSize(24)
	c.SetClip(newImg.Bounds())
	c.SetDst(newImg)
	c.SetSrc(image.NewUniform(color.Black))

	pt := freetype.Pt(15, newImg.Rect.Dy()-50)
	_, err = c.DrawString("Hello, World!", pt)
	if err != nil {
		panic(err)
	}
}

func drawRect(img image.Image) *image.RGBA {
	width := img.Bounds().Dx()
	height := img.Bounds().Dy()

	newImg := image.NewRGBA(image.Rect(0, 0, width, height))

	rectW := int(math.Round(float64(width) * 0.3))
	rectH := 50
	padding := 15
	draw.Draw(newImg, newImg.Bounds(), img, image.Point{0, 0}, draw.Over)

	rectImg := image.NewRGBA(image.Rect(0, 0, rectW, rectH))
	fillColor := color.RGBA{0, 0, 0, 64}
	draw.Draw(rectImg, rectImg.Bounds(), &image.Uniform{fillColor}, image.Point{0, 0}, draw.Src)

	rectPos := image.Point{0 + padding, (height - rectH) - padding}
	draw.Draw(newImg, image.Rect(rectPos.X, rectPos.Y, rectPos.X+rectW, rectPos.Y+rectH), rectImg, image.Point{0, 0}, draw.Over)
	return newImg
}
