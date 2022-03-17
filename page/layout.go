package page

import (
	glo "gioui.org/layout"
	"gioui.org/unit"
)

type (
	C = glo.Context
	D = glo.Dimensions
)

type DetailRow struct {
	PrimaryWidth float32
	glo.Inset
}

var DefaultInset = glo.UniformInset(unit.Dp(8))

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
