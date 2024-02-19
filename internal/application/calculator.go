package application

import (
	"github.com/diamondburned/gotk4/pkg/gtk/v4"
	"github.com/michelmb-dev/calculator-gtk/internal/ui"
)

type Tcalculator struct {
	*gtk.Application
	Window *gtk.ApplicationWindow
	Ui     *ui.Tui
}

func NewCalculator(app *gtk.Application) *Tcalculator {
	c := Tcalculator{Application: app}

	// create UI elements
	c.Ui = ui.CreateUiElements()

	// Create main window application
	c.Window = gtk.NewApplicationWindow(app)
	c.Window.SetTitle("Calculator")
	c.Window.SetTitlebar(c.Ui.Header)
	c.Window.SetChild(c.Ui)
	c.Window.SetResizable(false)

	return &c
}

func (c *Tcalculator) Show() {
	c.Window.Show()
}
