package gamelogic

import (
	cr "github.com/Oleg-MBO/blind_deity/creatures"
)

// Ground represent ground of game with all fields
type Ground struct {
	maxH, maxW int
	places     [][]cr.InhabitInterface
}

// NewGround create Ground
func NewGround(maxH, maxW int) *Ground {
	g := new(Ground)

	g.maxH = maxH
	g.maxW = maxW

	placesHArr := make([][]cr.InhabitInterface, maxH+1, maxH+1)
	for i := range placesHArr {
		placesHArr[i] = make([]cr.InhabitInterface, maxW+1, maxW+1)
	}
	g.places = placesHArr
	return g
}

// GetLimits return max height and wigth of fields
func (g *Ground) GetLimits() (h, w int) {
	return g.maxH, g.maxW
}

// IsInhabitExistOn return true if Inhabitexist on this field
func (g *Ground) IsInhabitExistOn(h, w int) bool {
	return g.places[h][w] != nil
}

// GetCreatureOn return Inhabit on h, w field or nil
func (g *Ground) GetCreatureOn(h, w int) cr.InhabitInterface {
	return g.places[h][w]
}

// SetCreatureOn put creature on h,w
func (g *Ground) SetCreatureOn(h, w int, inh cr.InhabitInterface) {
	g.places[h][w] = inh
}
