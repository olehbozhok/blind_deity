package basegui

import (
	gamelogic "github.com/Oleg-MBO/blind_deity/gameLogic"
	"github.com/Oleg-MBO/blind_deity/utils"
	"github.com/faiface/pixel"
)

type GroundGui struct {
	gamelogic.Ground
	hPx, wPx int

	pxField int
	sprite  *pixel.Sprite // sprite of image ground
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

func (g *GroundGui) Draw(t pixel.Target, matrix pixel.Matrix) {
	g.sprite.Draw(t, matrix)
	// TODO: Draw creatures
	h, w := g.GetLimits()
	for vh := 0; vh <= h; vh++ {
		for vw := 0; vw <= w; vw++ {
			cr := g.GetCreatureOn(vh, vw)
			if cr != nil {
				crPx := cr.GetPix()
				leftBottomMarix := matrix.Moved(pixel.V(float64(-g.wPx/2+crPx/2), float64(-g.hPx/2+crPx/2)))
				cr.Draw(t, leftBottomMarix.Moved(pixel.V(float64(vw*crPx), float64(vh*crPx))))
			}

		}
	}
}



