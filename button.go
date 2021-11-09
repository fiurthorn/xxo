package xxo

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
)

type Button struct {
	w, h int
	x, y int
}

func NewButton(x, y int, w, h int) *Button {
	return nil
}

func (g *Game) DrawNewGameButton(screen *ebiten.Image) {
	sw, sh := screen.Size()
	bw, bh := g.board.boardImage.Size()
	x, y := (sw-bw)/2, (sh-bh)/2+bh+10

	button := ebiten.NewImage(bw, tileSize/2)
	txt := "New Game"
	bound, _ := font.BoundString(f, txt)
	w := (bound.Max.X - bound.Min.X).Ceil()
	h := (bound.Max.Y - bound.Min.Y).Ceil()
	xT := (bw - w) / 2
	yT := (tileSize / 2) - h/2 + 3

	if g.board.Won() || g.board.Remaining() == 0 {
		button.Fill(activeColor)
		text.Draw(button, txt, f, xT, yT, tileBackgroundColor)
	} else {
		button.Fill(frameColor)
		text.Draw(button, txt, f, xT, yT, tileColor)
	}

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(x), float64(y))
	screen.DrawImage(button, op)
}
