package baseinhabitant

import (
	"fmt"
	"math/rand"

	cr "github.com/Oleg-MBO/blind_deity/creatures"
)

// NextStep return relative next position where Inhabitant want to be
func (i *BaseInhabitant) NextStep(relWatcher cr.RelativeWatcher) (x, y int) {
	// iterate over fields to find enemy
	for rH := 2; rH >= -2; rH-- {
		for rW := 2; rW >= -2; rW-- {
			if rH == 0 && rW == 0 {
				continue
			}
			if i.IsEnemy(relWatcher(rH, rW)) {
				// found enemy, trying to find safe field

				fmt.Printf("enemy on: h:%d , w:%d\n", rH, rW)

				isHPositive := rH >= 0
				isWPositive := rW >= 0
				// fmt.Printf("isHPositive:%t , isWPositive:%t \n", isHPositive, isWPositive)
				// fmt.Printf("i.IsSafeField(relWather, 1, 0):%t \n", i.IsSafeField(relWather, 1, 0))
				// fmt.Printf("i.IsSafeField(relWather, -1, 0):%t \n", i.IsSafeField(relWather, -1, 0))
				// fmt.Printf("i.IsSafeField(relWather, 0, 1):%t \n", i.IsSafeField(relWather, 0, 1))
				// fmt.Printf("i.IsSafeField(relWather, 0, 1):%t \n", i.IsSafeField(relWather, 0, 1))

				if !isHPositive && i.IsSafeField(relWatcher, 1, 0) {
					fmt.Println(1)
					return 1, 0
				} else if isHPositive && i.IsSafeField(relWatcher, -1, 0) {
					fmt.Println(2)
					return -1, 0
				} else if !isWPositive && i.IsSafeField(relWatcher, 0, 1) {
					fmt.Println(3)
					return 0, 1
				} else if isWPositive && i.IsSafeField(relWatcher, 0, -1) {
					fmt.Println(4)
					return 0, -1
				} else {
					fmt.Println(5)
					return 0, 0
				}

			}

		}
	}
	// if i.IsEnemy(relWather(1, 0)) {
	// 	return -1, 0
	// }
	// if i.IsEnemy(relWather(-1, 0)) {
	// 	return 1, 0
	// }
	// if i.IsEnemy(relWather(0, 1)) {
	// 	return 0, -1
	// }
	// if i.IsEnemy(relWather(0, -1)) {
	// 	return 0, 1
	// }
	return rand.Intn(i.maxMove+i.maxMove+1) - i.maxMove, rand.Intn(i.maxMove+i.maxMove+1) - i.maxMove
}

// IsSafeField return true if rH, rW is safe field
func (i *BaseInhabitant) IsSafeField(relWather cr.RelativeWatcher, rH, rW int) bool {
	if i.IsEnemy(relWather(rH+1, 0)) ||
		i.IsEnemy(relWather(rW-1, 0)) ||
		i.IsEnemy(relWather(0, rW+1)) ||
		i.IsEnemy(relWather(0, rW-1)) {
		return false
	}
	return true
}
