package img

import (
	"bytes"
	"errors"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"image/png"
	"io"
	"math"
	"os"
	"strings"

	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
)

func SaveImage(img *image.RGBA, filepath string) error {
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	return png.Encode(out, img)
}

func LoadImage(reader io.Reader) (image.Image, error) {
	img, _, err := image.Decode(reader)
	if err != nil {
		return nil, err
	}
	return img, nil
}

func ProcessImage(img image.Image, format string) (io.Reader, error) {
	var texto = "Hello Teste Teste Teste Teste Teste Teste Teste Teste Teste !"
	if len(texto) > 110 {
		return nil, errors.New("maximo de caracteres")
	}
	fontBytes, err := os.ReadFile("assets/arial.ttf")
	if err != nil {
		return nil, err
	}
	rsize, err := measureStringSize(img, texto, fontBytes)
	if err != nil {
		return nil, err
	}

	newImg := drawRect(img, rsize)
	err = drawText(newImg, texto, fontBytes)
	if err != nil {
		return nil, err
	}

	buf := new(bytes.Buffer)
	switch strings.ToLower(format) {
	case ".png":
		err = png.Encode(buf, newImg)
	case ".jpeg":
	case ".jpg":
		err = jpeg.Encode(buf, newImg, nil)
	default:
		return nil, fmt.Errorf("unsupported format: %s", format)
	}
	if err != nil {
		return nil, err
	}

	return buf, nil
}

func measureStringSize(img image.Image, text string, fontBytes []byte) (fixed.Int26_6, error) {
	f, err := truetype.Parse(fontBytes)
	if err != nil {
		return 0, err
	}

	width := img.Bounds().Dx()
	height := img.Bounds().Dy()

	newImg := image.NewRGBA(image.Rect(0, 0, width, height))

	ff := truetype.NewFace(f, &truetype.Options{
		Size: 24,
		DPI:  72,
	})

	drawer := font.Drawer{
		Dst:  newImg,
		Src:  newImg,
		Face: ff,
		Dot:  fixed.Point26_6{},
	}
	size := drawer.MeasureString(text)
	return size, nil
}

func drawText(newImg *image.RGBA, text string, fontBytes []byte) error {
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
	_, err = c.DrawString(text, pt)

	if err != nil {
		return err
	}
	return nil
}

func drawRect(img image.Image, size fixed.Int26_6) *image.RGBA {
	width := img.Bounds().Dx()
	height := img.Bounds().Dy()

	newImg := image.NewRGBA(image.Rect(0, 0, width, height))
	draw.Draw(newImg, newImg.Bounds(), img, image.Point{0, 0}, draw.Over)

	wpercent := (float64(width) * 0.3)
	println(size)
	println(size.Ceil())
	rectW := int(math.Max(wpercent, float64(size.Ceil()+int(float64(width)*0.1))))
	rectH := 60
	padding := 15
	rectPos := image.Rect(padding, height-rectH-padding, padding+rectW, height-padding)

	fillColor := color.RGBA{0, 0, 0, 64}
	draw.Draw(newImg, rectPos, &image.Uniform{fillColor}, image.Point{}, draw.Over)

	return newImg
}
