package main

import "github.com/faiface/pixel"

type Scene struct {
	w, h   int
	sprite *pixel.Sprite
	ground *Ground
}

func NewScene(width, height, pxField, numCreatures int) *Scene {
	gr := NewGround(width, height, pxField, numCreatures)
	scene := &Scene{w: width, h: height, ground: gr}
	return scene
}

func (s *Scene) Draw(t pixel.Target, matrix pixel.Matrix) {
	s.ground.Draw(t, matrix)
}

func (s *Scene) Update() {
	//TODO: handle change
}
