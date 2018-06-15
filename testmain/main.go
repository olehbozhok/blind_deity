package main

import (
	"math/rand"
	"time"

	_ "image/jpeg"
	_ "image/png"

	"github.com/faiface/pixel"

	"github.com/faiface/pixel/pixelgl"
	// "github.com/llgcode/draw2d/draw2dimg"
	"golang.org/x/image/colornames"

	cr "github.com/Oleg-MBO/blind_deity/creatures"
)

const (
	width      = 700
	height     = 700
	seed       = 6502
	numCircles = 4
)

func init() {
	rand.Seed(time.Now().Unix())
}

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "Blind Deily",
		Bounds: pixel.R(0, 0, float64(width), float64(height)),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}
	// win.Clear(colornames.Black)
	win.Clear(colornames.Brown)

	//  pixel.IM standart matrix
	IMCenter := pixel.IM.Moved(win.Bounds().Center())
	fieldSize := 10
	fieldSize = 300
	// countCreatures := 4

	cre := cr.NewBaseInhabitant(1, fieldSize)
	creImage := cre.GenImage()

	indPicCr := pixel.PictureDataFromImage(creImage)
	spriteCr := pixel.NewSprite(indPicCr, indPicCr.Bounds())
	// spriteCr.Draw(win, IMCenter)
	_ = spriteCr

	cre.Draw(win, IMCenter)

	// im := image.NewRGBA(image.Rect(0, 0, fieldSize, fieldSize))

	// utils.Drawcircle(im, fieldSize/2, fieldSize/2, fieldSize/2, color.White)
	// indPic := pixel.PictureDataFromImage(creImage)
	// // maxVec := indPic.Bounds().Max

	// sprite := pixel.NewSprite(indPic, indPic.Bounds())
	// sprite.Draw(win, IMCenter)

	for !win.Closed() {
		win.Update()

		if win.JustPressed(pixelgl.KeyQ) {
			return
		}

	}
}

func main() {
	pixelgl.Run(run)
}
