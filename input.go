package xxo

import "github.com/hajimehoshi/ebiten/v2"

func NewInput() *Input {
	return &Input{}
}

type Input struct {
	reset      bool
	isPressed  bool
	isReleased bool
	x, y       int
}

func (i *Input) Update() {
	if !ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) && i.isPressed {
		i.isPressed = false
		i.x, i.y = ebiten.CursorPosition()
		i.isReleased = true
	}
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		i.isPressed = true
	}

	if ebiten.IsKeyPressed(ebiten.KeyR) {
		i.reset = true
	}
}
