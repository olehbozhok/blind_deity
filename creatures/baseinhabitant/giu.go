package baseinhabitant

import (
	"image"

	"github.com/Oleg-MBO/blind_deity/utils"
	"github.com/faiface/pixel"
)

// Bulk set BaseInhabitant is bulk so need create new sprite
func (i *BaseInhabitant) Bulk() {
	i.bulk = true
}

// GetPix show how many pixels take cell
func (i *BaseInhabitant) GetPix() int {
	return i.pxPerson
}

// GenImage generate image of the Inhabitant
func (i *BaseInhabitant) GenImage() *image.RGBA {
	im := image.NewRGBA(image.Rect(0, 0, i.pxPerson, i.pxPerson))

	// draw.Draw(m, m.Bounds(), &image.Uniform{color.Black}, image.ZP, draw.Src)

	// c := color.RGBA{1, 2, 255, 255}
	// draw.Draw(im, im.Bounds(), &image.Uniform{c}, image.ZP, draw.Src)
	utils.Drawcircle(im, i.pxPerson/2, i.pxPerson/2, i.pxPerson/2, i.color)
	return im
}

// GetNewSprite return new sprite of Inhabitant
func (i *BaseInhabitant) GetNewSprite() *pixel.Sprite {
	i.bulk = false
	img := i.GenImage()
	indPic := pixel.PictureDataFromImage(img)
	i.sprite = pixel.NewSprite(indPic, indPic.Bounds())

	return i.sprite
}

// GetSprite return sprite of Inhabitant
func (i *BaseInhabitant) GetSprite() *pixel.Sprite {
	if i.bulk {
		return i.GetNewSprite()
	}
	return i.sprite
}
