package page

import (
	"gioui.org/layout"
	"gioui.org/x/component"
)

type Page interface {
	Actions() []component.AppBarAction
	Overflow() []component.OverflowAction
	Layout(gtx layout.Context) layout.Dimensions
	NavItem() component.NavItem
	Show()
}
