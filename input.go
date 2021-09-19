package xxo

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

func NewInput() *Input {
	return &Input{}
}

type Input struct {
	reset      bool
	isPressed  bool
	isReleased bool
	x, y       int
	touchId    ebiten.TouchID
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

	if inpututil.IsTouchJustReleased(i.touchId) {
		i.isReleased = true
	}

	touches := inpututil.JustPressedTouchIDs()
	if len(touches) > 0 {
		i.touchId = touches[0]
		i.x, i.y = ebiten.TouchPosition(touches[0])
	}

	if ebiten.IsKeyPressed(ebiten.KeyR) {
		i.reset = true
	}
}
