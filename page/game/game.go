package game

import (
	"image"
	"image/color"

	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"gioui.org/x/component"
	"gioui.org/x/outlay"

	"github.com/fiurthorn/xxo/lib"
	"github.com/fiurthorn/xxo/page"
	"github.com/fiurthorn/xxo/xxo"
)

type (
	C = layout.Context
	D = layout.Dimensions
)

type Page struct {
	*page.Router

	game       *xxo.Game
	grid       outlay.Grid
	newGameBtn widget.Clickable

	cells [9]*widget.Clickable
}

func New(router *page.Router) *Page {
	cells := [9]*widget.Clickable{}
	for i, length := 0, len(cells); i < length; i++ {
		cell := widget.Clickable{}
		cells[i] = &cell
	}
	return &Page{
		Router: router,
		game:   xxo.NewGame(),
		grid: outlay.Grid{
			Num:       3,
			Axis:      layout.Vertical,
			Alignment: layout.Middle,
		},
		cells: cells,
	}
}

func (p *Page) Actions() []component.AppBarAction {
	return []component.AppBarAction{
		component.SimpleIconAction(&p.newGameBtn, lib.NewGame,
			component.OverflowAction{
				Name: "New",
				Tag:  &p.newGameBtn,
			},
		),
	}
}

func (p *Page) Overflow() []component.OverflowAction {
	return []component.OverflowAction{}
}

func (p *Page) NavItem() component.NavItem {
	return component.NavItem{
		Name: "Game",
		Icon: lib.GameIcon,
	}
}

func (p *Page) Layout(gtx C) D {
	if p.newGameBtn.Clicked() {
		p.game.Board.ResetBoard()
	}

	return layout.Center.Layout(
		gtx,
		p.fill,
	)
}

func (p *Page) fill(gtx C) D {
	space := unit.Dp(8)
	line, ok := p.game.Board.Winning()

	return p.grid.Layout(gtx, 9, func(gtx layout.Context, i int) D {
		if p.cells[i].Clicked() && p.game.Board.IsEmpty(i) {
			p.game.Board.Set(i, p.game.Board.GetCurrent())

			// if p.game.Board.Stopped() {
			go func() {
				idx := p.game.BestMove(p.game.Board.GetCurrent())
				p.game.Board.Set(idx, p.game.Board.GetCurrent())
				// gtx.Ops.InvalidateOp()
			}()
			// }
		}

		size := lib.Max(
			lib.Min(gtx.Constraints.Min.X, gtx.Constraints.Min.Y),
			lib.Min(gtx.Constraints.Max.X, gtx.Constraints.Max.Y),
		)
		size = size / 7

		highlight := ok && p.game.Board.Contains(line, i)
		return layout.
			Inset{Top: space, Bottom: space, Left: space, Right: space}.
			Layout(gtx, p.button(gtx, i, size, highlight))
	})
}

func (p *Page) button(gtx C, idx int, size int, highlight bool) func(gtx C) D {
	return func(gtx C) D {
		gtx.Constraints = layout.Exact(image.Point{X: size, Y: size})

		clickable := p.cells[idx]
		player := p.game.Board.Get(idx).Symbol()

		btn := material.Button(p.Theme, clickable, player)
		btn.TextSize.V = float32(size / 2)
		if highlight {
			btn.Background = color.NRGBA{A: 0xff, R: 255, G: 64, B: 64}
		}

		d := btn.Layout(gtx)
		max := lib.Max(d.Size.X, d.Size.Y)

		return D{Size: image.Point{X: max, Y: max}}
	}
}
