package main

import (
	"image"
	"os"

	"image/color"
	"image/png"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	// "github.com/llgcode/draw2d/draw2dimg"
	"golang.org/x/image/colornames"
)

const (
	width      = 700
	height     = 700
	seed       = 6502
	numCircles = 4
)

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

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "Pixel Rocks!",
		Bounds: pixel.R(0, 0, float64(width), float64(height)),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}
	win.Clear(colornames.Black)

	//  pixel.IM standart matrix
	IMCenter := pixel.IM.Moved(win.Bounds().Center())
	fieldSize := 50

	fieldsImg := GenerateGameField(width, height, fieldSize)
	fieldPic := pixel.PictureDataFromImage(fieldsImg)
	fieldSprite := pixel.NewSprite(fieldPic, win.Bounds())

	pRect := win.Bounds()
	dest := image.NewRGBA(image.Rect(0, 0, width, height))
	// Drawcircle(dest, 50, 50, 50, colornames.Whitesmoke)
	for x := fieldSize / 2; x <= int(pRect.Max.X); x += fieldSize {
		Drawcircle(dest, x, fieldSize/2, fieldSize/2, colornames.Whitesmoke)
	}

	picData := pixel.PictureDataFromImage(dest)
	sprite2 := pixel.NewSprite(picData, win.Bounds())
	_ = sprite2
	// sprite2.Draw(win, IMCenter)
	for i := float64(0); i < 50; i++ {
		// sprite2.Draw(win, IMCenter.Moved(pixel.V(i, i)))
	}

	fieldSprite.Draw(win, IMCenter)

	sprite2.Draw(win, IMCenter)

	// win.Clear(colornames.Forestgreen)
	for !win.Closed() {
		win.Update()
		if win.JustPressed(pixelgl.KeyQ) {
			return
		}
		if win.JustPressed(pixelgl.KeyS) {
			w, _ := os.Create("blogmap.png")
			png.Encode(w, dest) //Encode writes the Image m to w in PNG format.
			w.Close()
		}
	}
}

func main() {
	pixelgl.Run(run)
}

