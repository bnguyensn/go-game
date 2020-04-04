package main

import (
	"fmt"

	"github.com/hajimehoshi/ebiten"
)

func update(screen *ebiten.Image) error {
	if ebiten.IsDrawingSkipped() {
		return nil
	}

	return nil
}

func main() {
	fmt.Println("Running...")

	if err := ebiten.Run(update, 320, 240, 2, "title"); err != nil {
		panic(err)
	}
}
