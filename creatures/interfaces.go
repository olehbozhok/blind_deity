package creatures

import "github.com/faiface/pixel"

// InhabitInterface interface of Inhabit
type InhabitInterface interface {
	NextStep() (x, y int)
	GetPix() int
	Draw(t pixel.Target, matrix pixel.Matrix)
}
