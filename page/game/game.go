package game

import (
	"image"

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
	config     *lib.Config
	grid       outlay.Grid
	newGameBtn widget.Clickable

	cells [9]*widget.Clickable

	auto chan int
}

func (p *Page) Show() {
	p.auto <- -1
}

func New(router *page.Router, config *lib.Config) *Page {
	cells := [9]*widget.Clickable{}
	for i, length := 0, len(cells); i < length; i++ {
		cell := widget.Clickable{}
		cells[i] = &cell
	}
	p := Page{
		Router: router,
		config: config,
		game:   xxo.NewGame(),
		grid: outlay.Grid{
			Num:       3,
			Axis:      layout.Horizontal,
			Alignment: layout.Middle,
		},
		cells: cells,
		auto:  make(chan int, 9),
	}
	go p.Auto()
	return &p
}

func (p *Page) Auto() {
	for task := range p.auto {
		p.game.Lock()
		if task >= 0 {
			p.game.Set(task, p.game.CurrentPlayer())
			p.game.TogglePlayer()
		}
		if !p.game.Stopped() {
			next := p.game.CurrentPlayer()
			if p.config.Ai1.Value && next == p.game.Player1() {
				p.auto <- p.game.BestMove(next)
			}
			if p.config.Ai2.Value && next == p.game.Player2() {
				p.auto <- p.game.BestMove(next)
			}
		}
		p.game.Unlock()
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
		p.game.ResetBoard()
		p.auto <- -1
	}

	return layout.Center.Layout(
		gtx,
		p.fill,
	)
}

func (p *Page) fill(gtx C) D {
	locked := p.game.Locked()

	space := unit.Dp(8)
	line, ok := p.game.Winning()
	ended := !locked && p.game.Stopped()

	return p.grid.Layout(gtx, 9, func(gtx layout.Context, i int) D {
		if p.cells[i].Clicked() && p.game.IsEmpty(i) {
			p.auto <- i
		}

		size := lib.Max(
			lib.Min(gtx.Constraints.Min.X, gtx.Constraints.Min.Y),
			lib.Min(gtx.Constraints.Max.X, gtx.Constraints.Max.Y),
		)
		size = size / 6

		highlight := !locked && ok && p.game.Contains(line, i)

		return layout.
			Inset{Top: space, Bottom: space, Left: space, Right: space}.
			Layout(gtx, p.button(gtx, i, size, highlight, ended))
	})
}

func (p *Page) button(gtx C, idx int, size int, highlight, ended bool) func(gtx C) D {
	return func(gtx C) D {
		gtx.Constraints = layout.Exact(image.Point{X: size, Y: size})

		clickable := p.cells[idx]
		player := p.game.Get(idx)

		btn := material.Button(p.Theme, clickable, player)
		btn.CornerRadius.U = unit.UnitPx
		btn.CornerRadius.V = float32(size / 2)
		btn.TextSize.U = unit.UnitPx
		btn.TextSize.V = float32(size / 2)
		if highlight {
			btn.Background = lib.SpanishOrange
		} else if ended {
			btn.Background = lib.Onyx
		}

		d := btn.Layout(gtx)
		max := lib.Max(d.Size.X, d.Size.Y)

		return D{Size: image.Point{X: max, Y: max}}
	}
}
