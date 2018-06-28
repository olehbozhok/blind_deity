package creatures

import (
	"github.com/Oleg-MBO/blind_deity/utils"
	"github.com/faiface/pixel"
)

// RelativeWatcher is used for see fields relative to h0,w0
// I don`t know in withc file put this type and not throw "import cycle not allowed"
type RelativeWatcher func(h, w int) InhabitInterface

// InhabitInterface interface of Inhabit
type InhabitInterface interface {
	NextStep(RelativeWatcher) (x, y int)
	IsBeget() (bool, utils.MoveVect, InhabitInterface)
	IsGoneAway() bool

	Force() int
	MakeHit(to InhabitInterface) (damage int)
	GotHit(damage int)

	GetPix() int
	GetSprite() *pixel.Sprite
}
