package controllers

import (
	"blind_deity/backend/components"
	"fmt"
)

func Init() {
	fmt.Println("init")
	i := components.NewInhabitant()
	g := components.NewGround(100, 100, 1)
	fmt.Println(i, g)
}
