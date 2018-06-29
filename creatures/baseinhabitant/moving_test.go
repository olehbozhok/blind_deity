package baseinhabitant

import (
	"testing"

	cr "github.com/Oleg-MBO/blind_deity/creatures"
	"github.com/Oleg-MBO/blind_deity/utils"
)

func TestBaseInhabitantNextStep(t *testing.T) {
	cr1 := NewBaseInhabitant(NewBaseInhabitantConf{
		MaxHealth:    20,
		MaxMove:      1,
		Fource:       1,
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
		t.Fail()
	}

	for rH := 2; rH >= -2; rH-- {
		for rW := 2; rW >= -2; rW-- {
			if rH == 0 && rW == 0 {
				continue
			}
			isRHPositive := rH > 0
			isRWPositive := rW > 0

			relWatcher := cr.RelativeWatcher(func(h, w int) cr.InhabitInterface {
				if h == rH && w == rW {
					return enemy
				}
				return nil
			})
			nextH, nextW := cr1.NextStep(relWatcher)
			isHPositive := nextH > 0
			isWPositive := nextW > 0

			// fmt.Printf("enemy on %d, %d; nextH=%d, nextW=%d\n", rH, rW, nextH, nextW)

			if isRHPositive == isHPositive && rH != 0 {
				t.Errorf("enemy on %d, %d; nextH=%d, nextW=%d isRHPositive must be %t ",
					rH, rW, nextH, nextW, !isRHPositive)
			}
			if isRWPositive == isWPositive && rW != 0 {
				t.Errorf("enemy on %d, %d; nextH=%d, nextW=%d isWPositive must be %t ",
					rH, rW, nextH, nextW, !isRWPositive)
			}

		}
	}

}
