package xxo

import "github.com/hajimehoshi/ebiten/v2"

type Widget interface {
	Draw(img *ebiten.Image)
}
