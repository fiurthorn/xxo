package settings

import (
	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"gioui.org/x/component"

	"github.com/fiurthorn/xxo/lib"
	"github.com/fiurthorn/xxo/page"
)

type (
	C = layout.Context
	D = layout.Dimensions
)

// Page holds the state for a page demonstrating the features of
// the AppBar component.
type Page struct {
	widget.List
	*page.Router

	config *lib.Config
}

// New constructs a Page with the provided router.
func New(router *page.Router, config *lib.Config) *Page {
	return &Page{
		Router: router,
		config: config,
	}
}

func (p *Page) Show() {}

func (p *Page) Actions() []component.AppBarAction {
	return []component.AppBarAction{}
}

func (p *Page) Overflow() []component.OverflowAction {
	return []component.OverflowAction{}
}

func (p *Page) NavItem() component.NavItem {
	return component.NavItem{
		Name: "Settings",
		Icon: lib.SettingsIcon,
	}
}

func (p *Page) Layout(gtx C) D {
	return layout.UniformInset(unit.Dp(10)).Layout(gtx, p.list)
}

func (p *Page) list(gtx C) D {
	return layout.Flex{
		Axis: layout.Vertical,
	}.Layout(gtx,
		layout.Flexed(1, material.Switch(p.Theme, &p.config.Ai1, "player1 AI").Layout),
		layout.Flexed(1, material.Switch(p.Theme, &p.config.Ai2, "player2 AI").Layout),
	)
}
