package img

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"math"
	"os"

	"github.com/golang/freetype"
)

func LoadImage(filePath string) (image.Image, error) {
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

func SaveImage(filePath string, img *image.RGBA) error {
	out, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer out.Close()

	return png.Encode(out, img)
}

func ProcessImage(img image.Image) (*image.RGBA, error) {
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
	rectPos := image.Rect(padding, height-rectH-padding, padding+rectW, height-padding)

	fillColor := color.RGBA{0, 0, 0, 64}
	draw.Draw(newImg, rectPos, &image.Uniform{fillColor}, image.Point{}, draw.Over)

	return newImg
}
