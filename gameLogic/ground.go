package gamelogic

import (
	cr "github.com/Oleg-MBO/blind_deity/creatures"
)

type Ground struct {
	maxH, maxW int
	places     []cr.InhabitInterface
}

func NewGround(maxH, maxW int) *Ground {
	g := new(Ground)

	g.maxH = maxH
	g.maxW = maxW

	places := make([]cr.InhabitInterface, maxH*maxW+maxW+1, maxH*maxW+maxW+1)
	g.places = places

	return g
}

func (g *Ground) GetLimits() (h, w int) {
	return g.maxH, g.maxW
}

func (g *Ground) IsInhabitExistOn(h, w int) bool {
	return g.places[h*w+w] != nil
}

func (g *Ground) IsCreatureOn(h, w int) cr.InhabitInterface {
	return g.places[h*w+w]
}

func (g *Ground) GetCreatureOn(h, w int) cr.InhabitInterface {
	return g.places[h*w+w]
}

func (g *Ground) SetCreatureOn(h, w int, inh cr.InhabitInterface) {
	g.places[h*w+w] = inh
}
