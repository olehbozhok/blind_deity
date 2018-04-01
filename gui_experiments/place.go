package main

type Place struct {
	x, y     int
	creature *Inhabitant
	// TODO: add params of field
}

func NewPlace(x, y int) Place {
	return Place{y: y, x: x}
}

func (p *Place) GetPosition() (int, int) {
	return p.x, p.y
}

// PopCreature delete creature from this field and return it
func (p *Place) getCreature() *Inhabitant {
	return p.creature

}

// PopCreature delete creature from this field and return it
func (p *Place) PopCreature() *Inhabitant {
	c := p.creature
	p.creature = nil
	return c
}

func (p *Place) SetCreature(inh *Inhabitant) bool {
	if p.creature != nil {
		return false
	}
	p.creature = inh
	return true
}
