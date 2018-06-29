package baseinhabitant

import (
	"image/color"
	"math/rand"

	"github.com/Oleg-MBO/blind_deity/utils"
	"github.com/faiface/pixel"

	cr "github.com/Oleg-MBO/blind_deity/creatures"
)

// BaseInhabitant represent base Inhabitant type
type BaseInhabitant struct {
	maxHealth  int
	currHealth int
	maxMove    int
	fource     int
	days       int

	percentBeget int
	percentDie   int

	pxPerson int
	color    color.Color
	sprite   *pixel.Sprite
	bulk     bool
}

// NewBaseInhabitantConf config to NewBaseInhabitant
type NewBaseInhabitantConf struct {
	MaxHealth int
	MaxMove   int
	Fource    int

	PercentBeget int
	PercentDie   int

	PxPerson int
	Color    color.Color
}

// NewBaseInhabitant used to create *BaseInhabitant
func NewBaseInhabitant(c NewBaseInhabitantConf) *BaseInhabitant {
	return &BaseInhabitant{
		maxHealth:    c.MaxHealth,
		currHealth:   c.MaxHealth,
		fource:       c.Fource,
		maxMove:      c.MaxMove,
		pxPerson:     c.PxPerson,
		percentBeget: c.PercentBeget,
		percentDie:   c.PercentDie,
		color:        c.Color,
		bulk:         true,
	}
}

// IsBeget when Inhabitant is beget return true, where and inhabit
func (i *BaseInhabitant) IsBeget() (bool, utils.MoveVect, cr.InhabitInterface) {
	if rand.Intn(100) <= i.percentBeget {

		mV := utils.MoveVect{}
		if rand.Intn(1) == 0 {
			mV.H = -1
		} else {
			mV.H = 1
		}
		if rand.Intn(1) == 0 {
			mV.W = -1
		} else {
			mV.W = 1
		}

		return true, mV, NewBaseInhabitant(NewBaseInhabitantConf{
			MaxMove:      i.maxMove,
			PxPerson:     i.pxPerson,
			Fource:       i.fource,
			PercentBeget: i.percentBeget,
			PercentDie:   i.percentDie,
			MaxHealth:    i.maxHealth,
			Color:        i.color,
		})
	}
	return false, utils.MoveVect{}, nil
}

// IsGoneAway true if Inhabitant is die
func (i *BaseInhabitant) IsGoneAway() bool {
	if i.currHealth <= 0 {
		return true
	}
	// return false
	return rand.Intn(100) <= i.percentDie
}

// Force return damage force of inhabitant
func (i *BaseInhabitant) Force() int {
	return i.fource
}

// IsEnemy return true if inh is enemy
func (i *BaseInhabitant) IsEnemy(inh cr.InhabitInterface) bool {
	if inh == nil {
		return false
	}

	if from1, ok := inh.(*BaseInhabitant); ok {
		if from1.maxMove == i.maxMove &&
			from1.fource == i.fource &&
			from1.percentBeget == i.percentBeget &&
			from1.percentDie == i.percentDie &&
			from1.maxHealth == i.maxHealth {
			return false
		}
	}
	return true
}

// MakeHit return damage force to hit inhabitant
func (i *BaseInhabitant) MakeHit(to cr.InhabitInterface) int {
	if !i.IsEnemy(to) {
		return 0
	}
	return i.Force()
}

// GotHit calls when inhabitant got hit and
// currHealth = currHealth - damage
func (i *BaseInhabitant) GotHit(damage int) {
	i.currHealth -= damage
}
