package main

import (
	"image"
	"image/color"
)

// GenerateGameField is drive game field
// with, height is a params of image
// pxSize is a size of field (drive don in the center of field)
func GenerateGameField(width, height, fieldSize int) *image.RGBA {
	m := image.NewRGBA(image.Rect(0, 0, width, height))

	// draw.Draw(m, m.Bounds(), &image.Uniform{color.Black}, image.ZP, draw.Src)

	for x := 0; x <= width; x += fieldSize {
		for y := 0; y <= height; y += fieldSize {
			m.Set(x-fieldSize/2, y-fieldSize/2, color.RGBA{150, 150, 250, 255})

			// if sum > threshold {
			// 	rect := image.Rect(x, y, x+pxSize, y+pxSize)
			// 	draw.Draw(m, rect, &image.Uniform{closestColor}, image.ZP, draw.Src)
			// } else {

			// }
		}
	}

	return m
}

func oilyPoint(image *image.RGBA, x, y int, col color.Color) {
	image.Set(x, y, col)
	image.Set(x-1, y, col)
	image.Set(x+1, y, col)
	image.Set(x, y+1, col)
	image.Set(x, y-1, col)
}

func Drawcircle(image *image.RGBA, x0, y0, radius int, col color.Color) {
	var x, y, dx, dy, derr int
	x = radius - 1
	y = 0
	dx = 1
	dy = 1
	derr = dx - (radius << 1)

	for x >= y {
		image.Set(x0+x, y0+y, col)
		image.Set(x0+y, y0+x, col)
		image.Set(x0-y, y0+x, col)
		image.Set(x0-x, y0+y, col)
		image.Set(x0-x, y0-y, col)
		image.Set(x0-y, y0-x, col)
		image.Set(x0+y, y0-x, col)
		image.Set(x0+x, y0-y, col)

		if derr < 0 {
			y++
			derr += dy
			dy += 2
		}

		if derr >= 0 {
			x--
			dx += 2
			derr += dx - (radius << 1)
		}
	}
}

// func DrawRectang(image *image.RGBA, x0, y0, x1, y1 int, col color.Color) {

// }

// func loadPicture(path string) (pixel.Picture, error) {
// 	file, err := os.Open(path)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer file.Close()
// 	img, _, err := image.Decode(file)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return pixel.PictureDataFromImage(img), nil
// }
