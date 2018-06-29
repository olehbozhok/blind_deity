package baseinhabitant

import (
	"image"
	"image/color"

	"github.com/Oleg-MBO/blind_deity/utils"
	"github.com/faiface/pixel"
)

type paramsSprite struct {
	// maxHealth  int
	// currHealth int
	// maxMove    int
	// fource     int
	radius int

	// percentBeget int
	// percentDie   int

	// pxPerson int
	color color.Color
}

type flyweight struct {
	cacheSprites map[paramsSprite]*pixel.Sprite
}

func newFlyweight() *flyweight {
	return &flyweight{make(map[paramsSprite]*pixel.Sprite)}
}

func (fl *flyweight) GetSprite(inh *BaseInhabitant) *pixel.Sprite {
	divider := 5 - (inh.days / 10)
	if divider < 2 {
		divider = 2
	}
	radius := inh.pxPerson / divider

	paramsSpr := paramsSprite{radius: radius, color: inh.color}
	sprite := fl.cacheSprites[paramsSpr]
	if sprite == nil {

		img := image.NewRGBA(image.Rect(0, 0, inh.pxPerson, inh.pxPerson))
		utils.Drawcircle(img, inh.pxPerson/2, inh.pxPerson/2, radius, inh.color)
		indPic := pixel.PictureDataFromImage(img)
		sprite = pixel.NewSprite(indPic, indPic.Bounds())
		fl.cacheSprites[paramsSpr] = sprite
	}
	return sprite
}
