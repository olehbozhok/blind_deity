package main

// import (
// 	"math/rand"
// 	"time"

// 	"github.com/faiface/pixel"
// 	"github.com/faiface/pixel/pixelgl"
// 	// "github.com/llgcode/draw2d/draw2dimg"
// 	"golang.org/x/image/colornames"
// )

// const (
// 	width      = 700
// 	height     = 700
// 	seed       = 6502
// 	numCircles = 4
// )

// func init() {
// 	rand.Seed(time.Now().Unix())
// }

// func run() {
// 	cfg := pixelgl.WindowConfig{
// 		Title:  "Blind Deily",
// 		Bounds: pixel.R(0, 0, float64(width), float64(height)),
// 		VSync:  true,
// 	}
// 	win, err := pixelgl.NewWindow(cfg)
// 	if err != nil {
// 		panic(err)
// 	}
// 	win.Clear(colornames.Black)

// 	//  pixel.IM standart matrix
// 	IMCenter := pixel.IM.Moved(win.Bounds().Center())
// 	fieldSize := 10

// 	gr := NewGround(width, height, fieldSize, 100)
// 	gr.Draw(win, IMCenter)

// 	// win.Clear(colornames.Forestgreen)
// 	for !win.Closed() {
// 		win.Update()
// 		gr.Draw(win, IMCenter)
// 		if win.JustPressed(pixelgl.KeyQ) {
// 			return
// 		}
// 	}
// }

// func main() {
// 	pixelgl.Run(run)
// }
