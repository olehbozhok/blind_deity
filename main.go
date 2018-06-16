package main

import (
	"math/rand"
	"time"

	"github.com/Oleg-MBO/blind_deity/basegui"
	"github.com/Oleg-MBO/blind_deity/utils"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	// "github.com/llgcode/draw2d/draw2dimg"
	"golang.org/x/image/colornames"

	cr "github.com/Oleg-MBO/blind_deity/creatures"
)

const (
	width  = 700
	height = 700
	// seed       = 6502
	// numCircles = 4
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
	win.Clear(colornames.Black)

	//  pixel.IM standart matrix
	IMCenter := pixel.IM.Moved(win.Bounds().Center())

	fieldSize := 10

	countCreatures := 50

	gr := basegui.NewGround(width, height, fieldSize)
	maxh, maxw := gr.GetLimits()
	for i := 0; i < countCreatures; i++ {
		randH := rand.Intn(maxh)
		randW := rand.Intn(maxw)

		// color := color.RGBA{
		// 	R: uint8(rand.Intn(255)), G: uint8(rand.Intn(255)), B: uint8(rand.Intn(255)), A: 255,
		// }
		cre := cr.NewBaseInhabitant(cr.NewBaseInhabitantConf{
			MaxHealth:    10,
			MaxMove:      1,
			Fource:       5,
			PercentBeget: 2,
			PercentDie:   1,

			PxPerson: fieldSize,
			Color:    utils.Green,
		})
		gr.SetCreatureOn(randH, randW, cre)
	}
	countCreatures = 100
	for i := 0; i < countCreatures; i++ {
		randH := rand.Intn(maxh)
		randW := rand.Intn(maxw)

		cre := cr.NewBaseInhabitant(cr.NewBaseInhabitantConf{
			MaxHealth:    6,
			MaxMove:      1,
			Fource:       2,
			PercentBeget: 2,
			PercentDie:   1,

			PxPerson: fieldSize,
			Color:    utils.Blue,
		})
		gr.SetCreatureOn(randH, randW, cre)
	}

	gr.Draw(win, IMCenter)

	// win.Clear(colornames.Forestgreen)
	ticker := time.NewTicker(100 * time.Millisecond)
	go func() {
		evSecond := time.NewTicker(100 * time.Millisecond)
		for range evSecond.C {
			gr.HandleNextStep()
		}
	}()

	for !win.Closed() {
		win.Update()

		if win.JustPressed(pixelgl.KeyQ) {
			return
		}

		select {
		case <-ticker.C:
			win.Clear(colornames.Black)
			gr.Draw(win, IMCenter)
		default:
		}
	}
}

func main() {
	pixelgl.Run(run)
}
