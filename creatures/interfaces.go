package creatures

import (
	"github.com/Oleg-MBO/blind_deity/utils"
	"github.com/faiface/pixel"
)

// InhabitInterface interface of Inhabit
type InhabitInterface interface {
	NextStep() (x, y int)
	IsBeget() (bool, utils.MoveVect, InhabitInterface)
	IsGoneAway() bool

	Force() int
	MakeHit(to InhabitInterface) (damage int)
	GotHit(damage int)

	GetPix() int
	GetSprite() *pixel.Sprite
}
