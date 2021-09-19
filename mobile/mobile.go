package mobile

import (
	"github.com/fiurthorn/xxo"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/mobile"
)

func init() {
	ebiten.SetWindowTitle("XXO")
	g := xxo.NewGame(xxo.NewInput(), xxo.NewBoard())
	mobile.SetGame(g)
}

func Dummy() {}
