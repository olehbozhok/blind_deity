package main

import (
	"math/rand"

	"github.com/Oleg-MBO/blind_deity/backend"
	"github.com/faiface/pixel"
)

type Ground struct {
	places       []*backend.Place
	maxCreatures int
	hPx, wPx     int
	hLen, wLen   int
	pxField      int
	sprite       *pixel.Sprite // sprite of image
}

func NewGround(w, h, pxField, maxCreatures int) *Ground {
	g := new(Ground)

	g.hPx = h
	g.wPx = w

	g.pxField = pxField

	hLen := h / pxField
	wLen := w / pxField

	g.hLen = hLen
	g.wLen = wLen

	// generate Places
	places := make([]*Place, hLen*wLen+1)
	for x := 0; x < hLen; x++ {
		for y := 0; y <= wLen; y++ {
			p := &Place{x: x, y: y}
			// cr := NewInhabitant(pxField)
			// p.SetCreature(cr)
			places[x*wLen+y] = p

		}
	}
	g.places = places

	for i := 1; i <= maxCreatures; i++ {
		// rEl := random(0, hLen*wLen-1)
		rEl := rand.Intn(hLen * wLen)
		cr := NewInhabitant(pxField)
		g.places[rEl].SetCreature(cr)
	}

	// create Sprite
	fieldsImg := GenerateGameField(w, h, pxField)
	fieldPic := pixel.PictureDataFromImage(fieldsImg)
	fieldSprite := pixel.NewSprite(fieldPic, pixel.R(0, 0, float64(w), float64(h)))
	g.sprite = fieldSprite

	return g
}

func (g *Ground) GetPlace(h, w int) *Place {
	return g.places[h*w+w]
}

func (g *Ground) PlaceExist(h, w int) bool {
	if h >= 0 && h <= g.hLen && w >= 0 && w <= g.wLen {
		return true
	}
	return false
}

// func (g *Ground) Update() {

// }

func (g *Ground) Draw(t pixel.Target, matrix pixel.Matrix) {
	g.sprite.Draw(t, matrix)
	// TODO: Draw creatures
	for _, place := range g.places {
		cr := place.GetCreature()
		if cr != nil {
			x, y := place.GetPosition()
			leftBottomMarix := matrix.Moved(pixel.V(float64(-g.wPx/2+cr.pxPerson/2), float64(-g.hPx/2+cr.pxPerson/2)))

			cr.Draw(t, leftBottomMarix.Moved(pixel.V(float64(x*g.pxField), float64(y*g.pxField))))
		}
	}

}
