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

	baseInh "github.com/Oleg-MBO/blind_deity/creatures/baseinhabitant"
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

	fieldSize := 10 * 2

	gr := basegui.NewGround(width, height, fieldSize)

	countCreatures := 10
	maxH, maxW := gr.GetLimits()
	// maxH, _ := gr.GetLimits()

	for i := 1; i < countCreatures; i++ {
		// randH := rand.Intn(maxH)
		// randW := rand.Intn(maxw)
		randH, randW := maxH-i, maxW-i
		// randH, randW := maxH-i, i

		// color := color.RGBA{
		// 	R: uint8(rand.Intn(255)), G: uint8(rand.Intn(255)), B: uint8(rand.Intn(255)), A: 255,
		// }
		cre := baseInh.NewBaseInhabitant(baseInh.NewBaseInhabitantConf{
			MaxHealth:    i,
			MaxMove:      0,
			Fource:       100,
			PercentBeget: -1,
			PercentDie:   -1,

			PxPerson: fieldSize,
			Color:    utils.Green,
		})
		gr.SetCreatureOn(randH, randW, cre)
	}

	// // countCreatures = 3
	// for i := 0; i < countCreatures; i++ {
	// 	randH := rand.Intn(maxH)
	// 	randW := rand.Intn(maxw)

	// 	cre := baseInh.NewBaseInhabitant(baseInh.NewBaseInhabitantConf{
	// 		MaxHealth:    2,
	// 		MaxMove:      1,
	// 		Fource:       80,
	// 		PercentBeget: -1,
	// 		PercentDie:   -1,

	// 		PxPerson: fieldSize,
	// 		Color:    utils.Blue,
	// 	})
	// 	gr.SetCreatureOn(randH, randW, cre)
	// }

	// cre := cr.NewBaseInhabitant(cr.NewBaseInhabitantConf{
	// 	MaxHealth:    50,
	// 	MaxMove:      1,
	// 	Fource:       60,
	// 	PercentBeget: 0,
	// 	PercentDie:   -1,

	// 	PxPerson: fieldSize,
	// 	Color:    utils.Green,
	// })
	// gr.SetCreatureOn(0, 0, cre)
	// gr.SetCreatureOn(-1, -0, cre)

	// gr.SetCreatureOn(-0, -0, cre)

	gr.Draw(win, IMCenter)

	// win.Clear(colornames.Forestgreen)
	ticker := time.NewTicker(250 * time.Millisecond)

	evSecond := time.NewTicker(250 * 4 / 4 * time.Millisecond)
	go func() {

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
