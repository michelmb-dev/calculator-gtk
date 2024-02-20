package main

import (
	_ "embed"
	"os"

	"github.com/diamondburned/gotk4/pkg/gio/v2"
	"github.com/diamondburned/gotk4/pkg/gtk/v4"
	"github.com/michelmb-dev/calculator-gtk/internal/application"
	"github.com/michelmb-dev/calculator-gtk/internal/gtkutil"
)

func main() {
	app := gtk.NewApplication("com.github.com.michelmb-dev.calculator-gtk", gio.ApplicationFlagsNone)
	app.ConnectActivate(func() {
		gtkutil.LoadCSS()
		calc := application.NewCalculator(app)
		calc.HandleConnectButtons()
		calc.HandleKeyboard()
		calc.Show()
	})

	if code := app.Run(os.Args); code > 0 {
		os.Exit(code)
	}
}
