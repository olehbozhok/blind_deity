package basegui

import (
	"fmt"

	gamelogic "github.com/Oleg-MBO/blind_deity/gameLogic"
	"github.com/Oleg-MBO/blind_deity/utils"
	"github.com/faiface/pixel"
)

type GroundGui struct {
	gamelogic.Ground
	// maxCreatures int
	// hLen, wLen   int
	hPx, wPx int

	pxField int
	sprite  *pixel.Sprite // sprite of image
}

func NewGround(hPx, wPx, pxField int) *GroundGui {
	hLen := hPx / pxField
	wLen := wPx / pxField

	// create Sprite
	fieldsImg := utils.GenerateGameField(wPx, hPx, pxField)
	fieldPic := pixel.PictureDataFromImage(fieldsImg)
	fieldSprite := pixel.NewSprite(fieldPic, pixel.R(0, 0, float64(wPx), float64(hPx)))

	g := &GroundGui{
		Ground:  *gamelogic.NewGround(hLen, wLen),
		hPx:     hPx,
		wPx:     hPx,
		pxField: pxField,
		sprite:  fieldSprite,
	}

	return g
}

// func (g *GroundGui) GetPlace(h, w int) *GroundGui {
// 	return g.places[h*w+w]
// }

// func (g *GroundGui) PlaceExist(h, w int) bool {
// 	if h >= 0 && h <= g.maxH && w >= 0 && w <= g.maxW {
// 		return true
// 	}
// 	return false
// }

// func (g *Ground) Update() {

// }

func (g *GroundGui) Draw(t pixel.Target, matrix pixel.Matrix) {
	g.sprite.Draw(t, matrix)
	// TODO: Draw creatures
	h, w := g.GetLimits()
	for vh := 0; vh <= h; vh++ {
		for vw := 0; vw <= w; vw++ {
			cr := g.GetCreatureOn(vh, vw)
			if cr != nil {
				fmt.Println(cr)
				fmt.Println(vh, vw)
				crPx := cr.GetPix()
				leftBottomMarix := matrix.Moved(pixel.V(float64(-g.wPx/2+crPx/2), float64(-g.hPx/2+crPx/2)))
				cr.Draw(t, leftBottomMarix.Moved(pixel.V(float64(vw*crPx), float64(vh*crPx))))
				// leftBottomMarix := matrix.Moved(pixel.V(float64(-g.wPx/2+crPx/2), float64(-g.hPx/2+crPx/2)))

				// cr.Draw(t, leftBottomMarix)

			}

		}
	}

	// for _, place := range g.places {
	// 	cr := place.GetCreature()
	// 	if cr != nil {
	// 		x, y := place.GetPosition()
	// 		leftBottomMarix := matrix.Moved(pixel.V(float64(-g.wPx/2+cr.pxPerson/2), float64(-g.hPx/2+cr.pxPerson/2)))

	// 		cr.Draw(t, leftBottomMarix.Moved(pixel.V(float64(x*g.pxField), float64(y*g.pxField))))
	// 	}
	// }
}
