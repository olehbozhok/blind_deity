package baseinhabitant

import (
	"fmt"
	"testing"

	cr "github.com/Oleg-MBO/blind_deity/creatures"
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

	relWatcher := cr.RelativeWatcher(func(h, w int) cr.InhabitInterface {
		if h == 1 && w == 0 {
			return enemy
		}
		return nil
	})
	h, w := cr1.NextStep(relWatcher)
	isHPositive := h > 0
	isWPositive := w > 0
	fmt.Printf("h:%d , w:%d\n", h, w)
	if h == 0 {
		t.Error("must be not 0")
	}

	if isHPositive {
		t.Error("h == 1 && w == 0 h must be not positive")
	}

	// if h != -1 && w != 0 {
	// 	t.Errorf("h must be -1 but is %d and w must be 0 but is %d", h, w)
	// }

	relWatcher = cr.RelativeWatcher(func(h, w int) cr.InhabitInterface {
		if h == -1 && w == 0 {
			return enemy
		}
		return nil
	})
	if !cr1.IsSafeField(relWatcher, 1, 0) {
		t.Log("i.IsSafeField(relWather, 1, 0) MUST BE TRUE") // <-- PROBLEM THIS
	}
	t.Logf("i.IsSafeField(relWather, -1, 0):%t \n", cr1.IsSafeField(relWatcher, -1, 0))
	t.Logf("i.IsSafeField(relWather, 0, 1):%t \n", cr1.IsSafeField(relWatcher, 0, 1))
	t.Logf("i.IsSafeField(relWather, 0, 1):%t \n", cr1.IsSafeField(relWatcher, 0, 1))
	h, w = cr1.NextStep(relWatcher)
	isHPositive = h >= 0
	isWPositive = w >= 0
	if h == 0 {
		t.Errorf("must be not 0 but is %d", h)
	}
	if !isHPositive {
		t.Error("h == -1 && w == 0 h must be  positive")
	}
	if isWPositive {

	}
	if h != 1 && w != 0 {
		t.Errorf("h must be 1 but is %d and w must be 0 but is %d", h, w)
	}

	relWatcher = cr.RelativeWatcher(func(h, w int) cr.InhabitInterface {
		if h == 0 && w == 1 {
			return enemy
		}
		return nil
	})
	h, w = cr1.NextStep(relWatcher)
	isHPositive = h > 0
	isWPositive = w > 0
	if !isHPositive {
		// t.Error("h == 1 && w == 0 h must be  positive")
	}
	if isWPositive {
		t.Error("h == 0 && w == 1 w must be not positive")

	}
	// if h != 0 && w != -1 {
	// 	t.Errorf("h must be 0 but is %d and w must be -1 but is %d", h, w)
	// }

	relWatcher = cr.RelativeWatcher(func(h, w int) cr.InhabitInterface {
		if h == 0 && w == -1 {
			return enemy
		}
		return nil
	})
	h, w = cr1.NextStep(relWatcher)
	isHPositive = h > 0
	isWPositive = w > 0
	if !isHPositive {
	}
	if isWPositive {
		t.Error("h == 0 && w == -1 w must be not positive")
	}
	// if h != 0 && w != 1 {
	// 	t.Errorf("h must be 0 but is %d and w must be 1 but is %d", h, w)
	// }

}
