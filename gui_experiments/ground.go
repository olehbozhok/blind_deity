package main

import "github.com/faiface/pixel"

type Ground struct {
	places       []Place
	maxCreatures int
	h, w         int
	hLen, wLen   int
	pxField      int
	sprite       *pixel.Sprite // sprite of image
}

func NewGround(w, h, pxField, maxCreatures int) *Ground {
	g := new(Ground)

	g.h = h
	g.w = w

	g.pxField = pxField

	hLen := h / pxField
	wLen := w / pxField

	g.hLen = hLen
	g.wLen = wLen

	// generate Places
	places := make([]Place, hLen*wLen)
	for x := 0; x <= hLen; x++ {
		for y := 0; y <= wLen; y++ {
			places[x*wLen+y] = Place{x: x, y: y}
		}

	}
	g.places = places

	// create Sprite
	// gfImag := GenerateGameField(w,h,pxField)

	fieldsImg := GenerateGameField(w, h, pxField)
	fieldPic := pixel.PictureDataFromImage(fieldsImg)
	fieldSprite := pixel.NewSprite(fieldPic, pixel.R(0, 0, float64(w), float64(h)))
	g.sprite = fieldSprite

	return g
}

func (g *Ground) Draw(t pixel.Target, matrix pixel.Matrix) {
	g.Draw(t, matrix)
	// TODO: Draw  
}
