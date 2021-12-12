package resizer

import (
	"bytes"
	"image"
	"image/jpeg"
	"image/png"

	"golang.org/x/image/draw"
)

type Resizer struct{}

func NewResizer() *Resizer {
	return &Resizer{}
}

func (r *Resizer) Resize(binary []byte) (*bytes.Buffer, error) {
	reader := bytes.NewReader(binary)
	img, t, err := image.Decode(reader)
	if err != nil {
		return nil, err
	}

	rct := img.Bounds()
	canvas := image.NewRGBA(image.Rect(0, 0, rct.Dx()*2, rct.Dy()*2))
	draw.CatmullRom.Scale(canvas, canvas.Bounds(), img, rct, draw.Over, nil)

	buff := bytes.NewBuffer([]byte{})

	switch t {
	case "jpeg":
		jpeg.Encode(buff, canvas, &jpeg.Options{Quality: 95})
	default:
		png.Encode(buff, canvas)
	}

	return buff, nil
}
