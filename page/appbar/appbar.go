package appbar

import (
	"image/color"
	"log"

	"gioui.org/layout"
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
	heartBtn, plusBtn, contextBtn          widget.Clickable
	exampleOverflowState, red, green, blue widget.Clickable
	bottomBar, customNavIcon               widget.Bool
	favorited                              bool
	widget.List
	*page.Router
}

// New constructs a Page with the provided router.
func New(router *page.Router) *Page {
	return &Page{
		Router: router,
	}
}

var _ page.Page = &Page{}

func (p *Page) Actions() []component.AppBarAction {
	return []component.AppBarAction{
		{
			OverflowAction: component.OverflowAction{
				Name: "Favorite",
				Tag:  &p.heartBtn,
			},
			Layout: func(gtx layout.Context, bg, fg color.NRGBA) layout.Dimensions {
				if p.heartBtn.Clicked() {
					p.favorited = !p.favorited
				}
				btn := component.SimpleIconButton(bg, fg, &p.heartBtn, lib.HeartIcon)
				btn.Background = bg
				if p.favorited {
					btn.Color = color.NRGBA{R: 200, A: 0xff}
				} else {
					btn.Color = fg
				}
				return btn.Layout(gtx)
			},
		},
		component.SimpleIconAction(&p.plusBtn, lib.PlusIcon,
			component.OverflowAction{
				Name: "Create",
				Tag:  &p.plusBtn,
			},
		),
	}
}

func (p *Page) Overflow() []component.OverflowAction {
	return []component.OverflowAction{
		{
			Name: "Example 1",
			Tag:  &p.exampleOverflowState,
		},
		{
			Name: "Example 2",
			Tag:  &p.exampleOverflowState,
		},
	}
}

func (p *Page) NavItem() component.NavItem {
	return component.NavItem{
		Name: "App Bar Features",
		Icon: lib.HomeIcon,
	}
}

const (
//settingNameColumnWidth    = .3
//settingDetailsColumnWidth = 1 - settingNameColumnWidth
)

func (p *Page) Layout(gtx C, th *material.Theme) D {
	p.List.Axis = layout.Vertical
	if p.plusBtn.Clicked() {
		log.Println("plus")
	}
	return material.List(th, &p.List).Layout(gtx, 1, func(gtx C, _ int) D {
		return layout.Flex{
			Alignment: layout.Middle,
			Axis:      layout.Vertical,
		}.Layout(gtx,
			layout.Rigid(func(gtx C) D {
				return page.DefaultInset.Layout(gtx, material.Body2(th, `The app bar widget provides a consistent interface element for triggering navigation and page-specific actions.

The controls below allow you to see the various features available in our App Bar implementation.`).Layout)
			}),
			layout.Rigid(func(gtx C) D {
				return page.DetailRow{}.Layout(gtx, material.Body1(th, "Contextual App Bar").Layout, func(gtx C) D {
					if p.contextBtn.Clicked() {
						p.Router.AppBar.SetContextualActions(
							[]component.AppBarAction{
								component.SimpleIconAction(&p.red, lib.HeartIcon,
									component.OverflowAction{
										Name: "House",
										Tag:  &p.red,
									},
								),
							},
							[]component.OverflowAction{
								{
									Name: "foo",
									Tag:  &p.blue,
								},
								{
									Name: "bar",
									Tag:  &p.green,
								},
							},
						)
						p.Router.AppBar.ToggleContextual(gtx.Now, "Contextual Title")
					}
					return material.Button(th, &p.contextBtn, "Trigger").Layout(gtx)
				})
			}),
			layout.Rigid(func(gtx C) D {
				return page.DetailRow{}.Layout(gtx,
					material.Body1(th, "Bottom App Bar").Layout,
					func(gtx C) D {
						if p.bottomBar.Changed() {
							if p.bottomBar.Value {
								p.Router.ModalNavDrawer.Anchor = component.Bottom
								p.Router.AppBar.Anchor = component.Bottom
							} else {
								p.Router.ModalNavDrawer.Anchor = component.Top
								p.Router.AppBar.Anchor = component.Top
							}
							p.Router.BottomBar = p.bottomBar.Value
						}

						return material.Switch(th, &p.bottomBar).Layout(gtx)
					})
			}),
			layout.Rigid(func(gtx C) D {
				return page.DetailRow{}.Layout(gtx,
					material.Body1(th, "Custom Navigation Icon").Layout,
					func(gtx C) D {
						if p.customNavIcon.Changed() {
							if p.customNavIcon.Value {
								p.Router.AppBar.NavigationIcon = lib.HomeIcon
							} else {
								p.Router.AppBar.NavigationIcon = lib.MenuIcon
							}
						}
						return material.Switch(th, &p.customNavIcon).Layout(gtx)
					})
			}),
			layout.Rigid(func(gtx C) D {
				return page.DetailRow{}.Layout(gtx,
					material.Body1(th, "Animated Resize").Layout,
					material.Body2(th, "Resize the width of your screen to see app bar actions collapse into or emerge from the overflow menu (as size permits).").Layout,
				)
			}),
			layout.Rigid(func(gtx C) D {
				return page.DetailRow{}.Layout(gtx,
					material.Body1(th, "Custom Action Buttons").Layout,
					material.Body2(th, "Click the heart action to see custom button behavior.").Layout)
			}),
		)
	})
}
