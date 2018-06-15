package main

import (
	"image"
	"image/color"

	"github.com/faiface/pixel"
)

type Inhabitant struct {
	health   int
	strong   int
	pxPerson int
	sprite   *pixel.Sprite
	bulk     bool
}

func NewInhabitant(pxPerson int) *Inhabitant {
	return &Inhabitant{pxPerson: pxPerson, bulk: true}
}

// GenImage generate image of the Inhabitant
func (i *Inhabitant) GenImage() *image.RGBA {
	im := image.NewRGBA(image.Rect(0, 0, i.pxPerson, i.pxPerson))
	// c := color.RGBA{0, 0, 255, 255}
	// draw.Draw(im, im.Bounds(), &image.Uniform{c}, image.ZP, draw.Src)
	Drawcircle(im, i.pxPerson/2, i.pxPerson/2, i.pxPerson/2, color.White)
	return im
}

func (i *Inhabitant) Bulk() {
	i.bulk = true
}

func (i *Inhabitant) Draw(t pixel.Target, matrix pixel.Matrix) {
	if i.bulk {
		i.bulk = false
		img := i.GenImage()
		indPic := pixel.PictureDataFromImage(img)
		// maxVec := indPic.Bounds().Max
		i.sprite = pixel.NewSprite(indPic, indPic.Bounds())
	}
	i.sprite.Draw(t, matrix)
}
