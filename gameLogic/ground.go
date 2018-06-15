package gamelogic

import (
	"sync"

	cr "github.com/Oleg-MBO/blind_deity/creatures"
)

// Ground represent ground of game with all fields
type Ground struct {
	maxH, maxW int
	places     [][]cr.InhabitInterface
	rwmutex    sync.RWMutex
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
	g.rwmutex.RLock()
	defer g.rwmutex.RUnlock()
	return g.places[h][w] != nil
}

// GetCreatureOn return Inhabit on h, w field or nil
func (g *Ground) GetCreatureOn(h, w int) cr.InhabitInterface {
	g.rwmutex.RLock()
	defer g.rwmutex.RUnlock()
	return g.places[h][w]
}

// SetCreatureOn put creature on h,w
func (g *Ground) SetCreatureOn(h, w int, inh cr.InhabitInterface) {
	g.rwmutex.Lock()
	defer g.rwmutex.Unlock()
	g.places[h][w] = inh
}

// HandleNextStep handle next step of life inhabittans on the ground
func (g *Ground) HandleNextStep() {
	g.rwmutex.Lock()
	defer g.rwmutex.Unlock()

	maxH, maxW := g.GetLimits()
	for vh := 0; vh <= maxH; vh++ {
		for vw := 0; vw <= maxW; vw++ {
			cr := g.places[vh][vw]
			if cr != nil {
				if cr.IsGoneAway() {
					g.places[vh][vw] = nil
					continue
				}
				toX, toY := cr.NextStep()
				toH := vh + toY
				toW := vw + toX
				if toH < 0 || toW < 0 || toH > maxH || toW > maxW {
					continue
				}
				if g.places[toH][toW] != nil {
					continue
				}

				g.places[vh][vw] = nil
				g.places[toH][toW] = cr

				isBeget, m, child := cr.IsBeget()
				if !isBeget {
					continue
				} else {
					toH = vh + m.H
					toW = vw + m.W
					if toH < 0 || toW < 0 || toH > maxH || toW > maxW || g.places[toH][toW] != nil {
						continue
					}
					g.places[toH][toW] = child
				}

			}

		}
	}
}
