package application

import (
	"github.com/diamondburned/gotk4/pkg/gtk/v4"
	"github.com/michelmb-dev/calculator-gtk/internal/ui"
)

type Tapplication struct {
	*gtk.Application
	Window     *gtk.ApplicationWindow
	Ui         *ui.Tui
	Calculator struct {
		Operand1  float64
		Operand2  float64
		Operator  string
		AwaitNext bool
	}
}

func NewCalculator(application *gtk.Application) *Tapplication {
	app := Tapplication{Application: application}

	// create UI elements
	app.Ui = ui.CreateUiElements()

	// Create main window application
	app.Window = gtk.NewApplicationWindow(application)
	app.Window.SetTitle("Calculator")
	app.Window.SetTitlebar(app.Ui.Header)
	app.Window.SetChild(app.Ui)
	app.Window.SetResizable(false)

	return &app
}

func (app *Tapplication) Show() {
	app.Window.Show()
}
