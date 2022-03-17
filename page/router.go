package page

import (
	"log"
	"time"

	"gioui.org/example/component/icon"
	"gioui.org/layout"
	"gioui.org/op/paint"
	"gioui.org/widget/material"
	"gioui.org/x/component"
)

type Router struct {
	Theme          *material.Theme
	pages          map[interface{}]Page
	current        interface{}
	NavAnim        component.VisibilityAnimation
	BottomBar      bool
	NonModalDrawer bool

	*component.ModalNavDrawer
	*component.AppBar
	*component.ModalLayer
}

func NewRouter(th *material.Theme) Router {
	modal := component.NewModal()

	nav := component.NewNav("XXO", "tic-tac-toe")
	modalNav := component.ModalNavFrom(&nav, modal)

	bar := component.NewAppBar(modal)
	bar.NavigationIcon = icon.MenuIcon

	na := component.VisibilityAnimation{
		State:    component.Invisible,
		Duration: 250 * time.Millisecond,
	}
	return Router{
		pages:          map[interface{}]Page{},
		Theme:          th,
		ModalLayer:     modal,
		ModalNavDrawer: modalNav,
		AppBar:         bar,
		NavAnim:        na,
		NonModalDrawer: false,
	}
}

func (r *Router) Register(tag interface{}, p Page) {
	r.pages[tag] = p
	navItem := p.NavItem()
	navItem.Tag = tag
	if r.current == interface{}(nil) {
		r.current = tag
		r.AppBar.Title = navItem.Name
		r.AppBar.SetActions(p.Actions(), p.Overflow())
	}
	r.ModalNavDrawer.AddNavItem(navItem)
}

func (r *Router) SwitchTo(tag interface{}) {
	p, ok := r.pages[tag]
	if !ok {
		return
	}
	navItem := p.NavItem()
	r.current = tag
	r.AppBar.Title = navItem.Name
	r.AppBar.SetActions(p.Actions(), p.Overflow())
}

func (r *Router) Layout(gtx layout.Context) layout.Dimensions {
	for _, event := range r.AppBar.Events(gtx) {
		switch event := event.(type) {
		case component.AppBarNavigationClicked:
			if r.NonModalDrawer {
				r.NavAnim.ToggleVisibility(gtx.Now)
			} else {
				r.ModalNavDrawer.Appear(gtx.Now)
				r.NavAnim.Disappear(gtx.Now)
			}
		case component.AppBarContextMenuDismissed:
			log.Printf("Context menu dismissed: %v", event)
		case component.AppBarOverflowActionClicked:
			log.Printf("Overflow action selected: %v", event)
		}
	}

	if r.ModalNavDrawer.NavDestinationChanged() {
		r.SwitchTo(r.ModalNavDrawer.CurrentNavDestination())
	}

	paint.Fill(gtx.Ops, r.Theme.Palette.Bg)

	content := layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
		return layout.Flex{}.Layout(gtx, r.layoutNavigation(), r.layoutPage())
	})

	bar := r.layoutAppbar()
	flex := layout.Flex{Axis: layout.Vertical}
	if r.BottomBar {
		flex.Layout(gtx, content, bar)
	} else {
		flex.Layout(gtx, bar, content)
	}
	r.ModalLayer.Layout(gtx, r.Theme)
	return layout.Dimensions{Size: gtx.Constraints.Max}
}

func (r *Router) layoutAppbar() layout.FlexChild {
	return layout.Rigid(func(gtx layout.Context) layout.Dimensions {
		return r.AppBar.Layout(gtx, r.Theme, "home", "more")
	})
}

func (r *Router) layoutNavigation() layout.FlexChild {
	return layout.Rigid(func(gtx layout.Context) layout.Dimensions {
		gtx.Constraints.Max.X /= 3
		return r.NavDrawer.Layout(gtx, r.Theme, &r.NavAnim)
	})
}

func (r *Router) layoutPage() layout.FlexChild {
	return layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
		return r.pages[r.current].Layout(gtx)
	})
}
