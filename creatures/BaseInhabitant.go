package creatures

import (
	"image"
	"image/color"
	"math/rand"

	"github.com/Oleg-MBO/blind_deity/utils"
	"github.com/faiface/pixel"
)

// BaseInhabitant represent base Inhabitant type
type BaseInhabitant struct {
	maxMove int

	pxPerson int
	sprite   *pixel.Sprite
	bulk     bool
}

// NewBaseInhabitant used to create *BaseInhabitant
func NewBaseInhabitant(maxMove, pxPerson int) *BaseInhabitant {
	return &BaseInhabitant{maxMove: maxMove, pxPerson: pxPerson, bulk: true}
}

// NextStep return relative next position where Inhabitant want to be
func (i *BaseInhabitant) NextStep() (x, y int) {
	return rand.Intn(i.maxMove+i.maxMove+1) - i.maxMove, rand.Intn(i.maxMove+i.maxMove+1) - i.maxMove
}

// GenImage generate image of the Inhabitant
func (i *BaseInhabitant) GenImage() *image.RGBA {
	im := image.NewRGBA(image.Rect(0, 0, i.pxPerson, i.pxPerson))

	// draw.Draw(m, m.Bounds(), &image.Uniform{color.Black}, image.ZP, draw.Src)

	// c := color.RGBA{1, 2, 255, 255}
	// draw.Draw(im, im.Bounds(), &image.Uniform{c}, image.ZP, draw.Src)
	utils.Drawcircle(im, i.pxPerson/2, i.pxPerson/2, i.pxPerson/2, color.White)
	return im
}

// Bulk set BaseInhabitant is bulk so need create new sprite
func (i *BaseInhabitant) Bulk() {
	i.bulk = true
}

// Getpix show how many pixels take cell
func (i *BaseInhabitant) GetPix() int {
	return i.pxPerson
}

// Draw used to draw Inhabitant
func (i *BaseInhabitant) Draw(t pixel.Target, matrix pixel.Matrix) {
	if i.bulk {
		i.bulk = false
		img := i.GenImage()
		indPic := pixel.PictureDataFromImage(img)
		// maxVec := indPic.Bounds().Max
		i.sprite = pixel.NewSprite(indPic, indPic.Bounds())
	}
	i.sprite.Draw(t, matrix)
}
