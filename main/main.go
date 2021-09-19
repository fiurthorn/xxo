package main

import (
	"log"

	"github.com/fiurthorn/xxo"
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	g := xxo.NewGame(xxo.NewInput(), xxo.NewBoard())

	ebiten.SetWindowSize(xxo.ScreenWidth, xxo.ScreenHeight)
	ebiten.SetWindowTitle("XXO")

	if err := ebiten.RunGame(g); err != nil {
		log.Println(err)
	}

}
