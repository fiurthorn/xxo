package main

import (
	"flag"
	"log"
	"os"

	"gioui.org/app"
	"gioui.org/font/gofont"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/unit"
	"gioui.org/widget/material"
	"github.com/fiurthorn/xxo/page"
	"github.com/fiurthorn/xxo/page/game"
	"github.com/fiurthorn/xxo/page/settings"
)

type (
	C = layout.Context
	D = layout.Dimensions
)

func main() {
	flag.Parse()
	go func() {
		w := app.NewWindow()
		if err := loop(w); err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()
	app.Main()
}

func loop(w *app.Window) error {
	th := material.NewTheme(gofont.Collection())
	th.TextSize = unit.Sp(14.)
	var ops op.Ops

	router := page.NewRouter(th)
	router.Register(0, game.New(&router))
	router.Register(1, settings.New(&router))

	for {
		select {
		case e := <-w.Events():
			switch e := e.(type) {
			case system.DestroyEvent:
				{
					return e.Err
				}
			case system.FrameEvent:
				{
					gtx := layout.NewContext(&ops, e)
					router.Layout(gtx)
					e.Frame(gtx.Ops)
				}
			case app.ConfigEvent:
				{
					// log.Println(e.Config.Size.X, e.Config.Size.Y)
				}
			}
		}
	}
}
