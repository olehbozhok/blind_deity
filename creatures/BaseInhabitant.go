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
	maxHealth  int
	currHealth int
	maxMove    int
	fource     int

	percentBeget int
	percentDie   int

	pxPerson int
	sprite   *pixel.Sprite
	bulk     bool
}

// NewBaseInhabitantConf config to NewBaseInhabitant
type NewBaseInhabitantConf struct {
	MaxHealth int
	MaxMove   int
	Fource    int

	PercentBeget int
	PercentDie   int

	PxPerson int
}

// NewBaseInhabitant used to create *BaseInhabitant
func NewBaseInhabitant(c NewBaseInhabitantConf) *BaseInhabitant {
	return &BaseInhabitant{
		maxHealth:    c.MaxHealth,
		currHealth:   c.MaxHealth,
		fource:       c.Fource,
		maxMove:      c.MaxMove,
		pxPerson:     c.PxPerson,
		percentBeget: c.PercentBeget,
		percentDie:   c.PercentDie,
		bulk:         true,
	}
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

// GetPix show how many pixels take cell
func (i *BaseInhabitant) GetPix() int {
	return i.pxPerson
}

// Draw used to draw Inhabitant
func (i *BaseInhabitant) Draw(t pixel.Target, matrix pixel.Matrix) {
	if i.bulk {
		i.bulk = false
		img := i.GenImage()
		indPic := pixel.PictureDataFromImage(img)
		i.sprite = pixel.NewSprite(indPic, indPic.Bounds())
	}
	i.sprite.Draw(t, matrix)
}

// IsBeget when Inhabitant is beget return true, where and inhabit
func (i *BaseInhabitant) IsBeget() (bool, utils.MoveVect, InhabitInterface) {
	if rand.Intn(100) <= i.percentBeget {

		mV := utils.MoveVect{}
		if rand.Intn(1) == 0 {
			mV.H = -1
		} else {
			mV.H = 1
		}
		if rand.Intn(1) == 0 {
			mV.W = -1
		} else {
			mV.W = 1
		}

		return true, mV, NewBaseInhabitant(NewBaseInhabitantConf{
			MaxMove:      i.maxMove,
			PxPerson:     i.pxPerson,
			Fource:       i.fource,
			PercentBeget: i.percentBeget,
			PercentDie:   i.percentDie,
			MaxHealth:    i.maxHealth,
		})
	}
	return false, utils.MoveVect{}, nil
}

// IsGoneAway true if Inhabitant is die
func (i *BaseInhabitant) IsGoneAway() bool {
	if i.currHealth <= 0 {
		return true
	}
	// return false
	return rand.Intn(100) <= i.percentDie
}

func (i *BaseInhabitant) Force() int {
	return i.fource
}

func (i *BaseInhabitant) GotHit(from InhabitInterface) {
	i.currHealth = i.currHealth - from.Force()
}
