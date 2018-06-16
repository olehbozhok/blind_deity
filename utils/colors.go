package utils

import (
	"image/color"
	"math/rand"
)

var (
	// https://developer.mozilla.org/ru/docs/Web/CSS/CSS_Colors/Color_picker_tool

	// White color RGBA
	White = color.RGBA{
		R: 255, G: 255, B: 255, A: 255,
	}

	// Green color RGBA
	Green = color.RGBA{
		R: 63, G: 191, B: 63, A: 255,
	}

	// Blue color RGBA
	Blue = color.RGBA{
		R: 63, G: 63, B: 191, A: 255,
	}
)

// RandColorRGBA return random RGBA color
func RandColorRGBA() color.RGBA {
	return color.RGBA{
		R: uint8(rand.Intn(255)), G: uint8(rand.Intn(255)), B: uint8(rand.Intn(255)), A: 255,
	}
}
