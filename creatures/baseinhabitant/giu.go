package baseinhabitant

import (
	"github.com/faiface/pixel"
)

var spriteCache = newFlyweight()

// GetPix show how many pixels take cell
func (i *BaseInhabitant) GetPix() int {
	return i.pxPerson
}

// GetSprite return sprite of Inhabitant
func (i *BaseInhabitant) GetSprite() *pixel.Sprite {

	return spriteCache.GetSprite(i)
}
