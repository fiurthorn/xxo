package page

import (
	glo "gioui.org/layout"
	"gioui.org/unit"
)

type (
	C = glo.Context
	D = glo.Dimensions
)

// DetailRow lays out two widgets in a horizontal row, with the left
// widget considered the "Primary" widget.
type DetailRow struct {
	// PrimaryWidth is the fraction of the available width that should
	// be allocated to the primary widget. It should be in the range
	// (0,1.0]. Defaults to 0.3 if not set.
	PrimaryWidth float32
	// Inset is automatically applied to both widgets. This inset is
	// required, and will default to a uniform 8DP inset if not set.
	glo.Inset
}

var DefaultInset = glo.UniformInset(unit.Dp(8))

// Layout the DetailRow with the provided widgets.
func (d DetailRow) Layout(gtx C, primary, detail glo.Widget) D {
	if d.PrimaryWidth == 0 {
		d.PrimaryWidth = 0.3
	}
	if d.Inset == (glo.Inset{}) {
		d.Inset = DefaultInset
	}
	return glo.Flex{Alignment: glo.Middle}.Layout(gtx,
		glo.Flexed(d.PrimaryWidth, func(gtx C) D {
			return d.Inset.Layout(gtx, primary)
		}),
		glo.Flexed(1-d.PrimaryWidth, func(gtx C) D {
			return d.Inset.Layout(gtx, detail)
		}),
	)
}
