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

// getCreatureOn hidden (without rwmutex.RLock())
func (g *Ground) getCreatureOn(h, w int) cr.InhabitInterface {
	maxH, maxW := g.GetLimits()
	if h < 0 {
		h = h + maxH + 1
	}
	if w < 0 {
		w = w + maxW + 1
	}
	if h >= maxH {
		h = h - maxH + 1
	}
	if w >= maxW {
		w = w - maxW + 1
	}
	return g.places[h][w]
}

// getRelativeWatcher hidden (without rwmutex.RLock())
// show fields relative to h0, w0
func (g *Ground) getRelativeWatcher(h0, w0 int) cr.RelativeWatcher {
	return cr.RelativeWatcher(func(h, w int) cr.InhabitInterface {
		return g.getCreatureOn(h+h0, w+w0)
	})
}

// GetCreatureOn return Inhabit on h, w field or nil
func (g *Ground) GetCreatureOn(h, w int) cr.InhabitInterface {
	g.rwmutex.RLock()
	defer g.rwmutex.RUnlock()
	return g.places[h][w]
}

func (g *Ground) setCreatureOn(h, w int, inh cr.InhabitInterface) {
	maxH, maxW := g.GetLimits()
	if h < 0 {
		h = h + maxH + 1
	}
	if w < 0 {
		w = w + maxW + 1

	}
	if h > maxH {
		h = h - maxH - 1
	}
	if w > maxW {
		w = w - maxW - 1
	}
	g.places[h][w] = inh
}

// SetCreatureOn put creature on h,w
func (g *Ground) SetCreatureOn(h, w int, inh cr.InhabitInterface) {
	g.rwmutex.Lock()
	defer g.rwmutex.Unlock()
	g.setCreatureOn(h, w, inh)
}

// HandleNextStep handle next step of life inhabittans on the ground
func (g *Ground) HandleNextStep() {
	g.rwmutex.Lock()
	defer g.rwmutex.Unlock()

	maxH, maxW := g.GetLimits()

	setMoveInhabbit := make(map[cr.InhabitInterface]bool)

	// MOVE Creature
	for vh := 0; vh <= maxH; vh++ {
		for vw := 0; vw <= maxW; vw++ {
			cr1 := g.places[vh][vw]
			if cr1 != nil {
				if setMoveInhabbit[cr1] {
					continue
				}
				if cr1.IsGoneAway() {
					g.setCreatureOn(vh, vw, nil)
					continue
				}

				setMoveInhabbit[cr1] = true

				toRH, toRW := cr1.NextStep(g.getRelativeWatcher(vh, vw))
				toH := vh + toRH
				toW := vw + toRW

				if g.getCreatureOn(toH, toW) == nil {
					g.setCreatureOn(vh, vw, nil)
					g.setCreatureOn(toH, toW, cr1)
				}
			}
		}
	}

	// Make hits
	for vh := 0; vh <= maxH; vh++ {
		for vw := 0; vw <= maxW; vw++ {
			cr1 := g.places[vh][vw]
			if cr1 != nil {
				if cr2 := g.getCreatureOn(vh+1, vw+1); cr2 != nil {
					cr1.GotHit(cr2.MakeHit(cr1))
				}
				if cr2 := g.getCreatureOn(vh-1, vw+1); cr2 != nil {
					cr1.GotHit(cr2.MakeHit(cr1))
				}
				if cr2 := g.getCreatureOn(vh+1, vw-1); cr2 != nil {
					cr1.GotHit(cr2.MakeHit(cr1))
				}
				if cr2 := g.getCreatureOn(vh-1, vw-1); cr2 != nil {
					cr1.GotHit(cr2.MakeHit(cr1))
				}
				if cr1.IsGoneAway() {
					g.setCreatureOn(vh, vw, nil)
				}
			}
		}
	}

	// creatures begets
	setBegetsInhabbit := make(map[cr.InhabitInterface]bool)

	for vh := 0; vh <= maxH; vh++ {
		for vw := 0; vw <= maxW; vw++ {
			cr1 := g.places[vh][vw]
			if cr1 != nil && !cr1.IsGoneAway() {
				if setBegetsInhabbit[cr1] {
					continue
				}
				isBeget, m, child := cr1.IsBeget()
				if isBeget {
					toH := vh + m.H
					toW := vw + m.W
					if g.getCreatureOn(toH, toW) == nil {
						g.setCreatureOn(toH, toW, child)
						setMoveInhabbit[child] = true

					}
				}
			}
		}
	}
}
