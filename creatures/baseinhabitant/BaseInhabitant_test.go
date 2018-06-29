package baseinhabitant

import (
	"testing"

	"github.com/Oleg-MBO/blind_deity/utils"
)

func TestBaseInhabitantIsEnemy(t *testing.T) {
	cr1 := NewBaseInhabitant(NewBaseInhabitantConf{
		MaxHealth:    20,
		MaxMove:      1,
		Fource:       1,
		PercentBeget: 9,
		PercentDie:   4,

		PxPerson: 10,
		Color:    utils.Green,
	})
	cr2 := NewBaseInhabitant(NewBaseInhabitantConf{
		MaxHealth:    21,
		MaxMove:      1,
		Fource:       1,
		PercentBeget: 9,
		PercentDie:   4,

		PxPerson: 10,
		Color:    utils.Green,
	})
	if !cr1.IsEnemy(cr2) {
		t.Error("cr1 and cr2 is Enemis")

		t.Fail()
	}
}

func TestBaseInhabitantIsGoneAway(t *testing.T) {
	cr1 := NewBaseInhabitant(NewBaseInhabitantConf{
		MaxHealth:    20,
		MaxMove:      1,
		Fource:       1,
		PercentBeget: 9,
		PercentDie:   4,

		PxPerson: 10,
		Color:    utils.Green,
	})
	cr1.GotHit(20)
	if !cr1.IsGoneAway() {
		t.Error("Creature must be die")
	}
}

func TestBaseInhabitantMakeHit(t *testing.T) {
	cr1 := NewBaseInhabitant(NewBaseInhabitantConf{
		MaxHealth:    20,
		MaxMove:      1,
		Fource:       21,
		PercentBeget: 9,
		PercentDie:   4,

		PxPerson: 10,
		Color:    utils.Green,
	})
	enemy := NewBaseInhabitant(NewBaseInhabitantConf{
		MaxHealth:    21,
		MaxMove:      1,
		Fource:       1,
		PercentBeget: 9,
		PercentDie:   4,

		PxPerson: 10,
		Color:    utils.Green,
	})
	if !cr1.IsEnemy(enemy) {
		t.Error("cr1 and cr2 is Enemis")
	}

	damage := cr1.MakeHit(enemy)
	if damage != 21 {
		t.Error("damage must be 21")
	}
}
