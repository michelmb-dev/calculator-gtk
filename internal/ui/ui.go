package ui

import (
	_ "embed"

	"github.com/diamondburned/gotk4/pkg/gtk/v4"
	"github.com/michelmb-dev/calculator-gtk/internal/gtkutil"
)

type Tbutton struct {
	*gtk.Button
	Label string
	Type  string // "operator" | "operand" | "result" | "reset"
}

type Tui struct {
	*gtk.Box
	Header       *gtk.HeaderBar
	DisplayEntry *gtk.Entry
	ButtonsGrid  *gtk.Grid
	Buttons      [][]Tbutton
}

//go:embed ui.css
var css string

func init() { gtkutil.AddCSS(css) }

func CreateUiElements() *Tui {
	ui := Tui{}

	// create header
	ui.Header = gtk.NewHeaderBar()

	// Create main Box
	ui.Box = gtk.NewBox(gtk.OrientationVertical, 10)
	ui.SetMarginTop(10)
	ui.SetMarginBottom(10)
	ui.SetMarginStart(10)
	ui.SetMarginEnd(10)

	// Create display entry
	ui.DisplayEntry = gtk.NewEntry()
	ui.DisplayEntry.SetCanFocus(false)
	ui.DisplayEntry.SetFocusable(false)
	ui.DisplayEntry.SetFocusOnClick(false)
	ui.DisplayEntry.SetEditable(false)
	ui.DisplayEntry.Delegate().SetAlignment(1)

	// Create grid buttons
	ui.ButtonsGrid = gtk.NewGrid()
	ui.ButtonsGrid.SetColumnHomogeneous(true)
	ui.ButtonsGrid.SetRowHomogeneous(true)
	ui.ButtonsGrid.SetRowSpacing(5)
	ui.ButtonsGrid.SetColumnSpacing(5)

	// Create operator and operand buttons
	ui.Buttons = *createCalculatorButtons(ui.ButtonsGrid)

	ui.Box.Append(ui.DisplayEntry)
	ui.Box.Append(ui.ButtonsGrid)

	return &ui
}

func createCalculatorButtons(buttonsGrid *gtk.Grid) *[][]Tbutton {
	buttons := [][]Tbutton{
		{
			{Label: "AC", Type: "reset"},
			{Label: "%", Type: "operator"},
			{Label: "√", Type: "operator"},
			{Label: "/", Type: "operator"},
		},
		{
			{Label: "7", Type: "operand"},
			{Label: "8", Type: "operand"},
			{Label: "9", Type: "operand"},
			{Label: "*", Type: "operator"},
		},

		{
			{Label: "4", Type: "operand"},
			{Label: "5", Type: "operand"},
			{Label: "6", Type: "operand"},
			{Label: "-", Type: "operator"},
		},

		{
			{Label: "1", Type: "operand"},
			{Label: "2", Type: "operand"},
			{Label: "3", Type: "operand"},
			{Label: "+", Type: "operator"},
		},

		{
			{Label: "0", Type: "operand"},
			{Label: ".", Type: "operand"},
			{Label: "π", Type: "operand"},
			{Label: "=", Type: "result"},
		},
	}

	for rowId, row := range buttons {
		for colId, btn := range row {
			button := gtk.NewButtonWithLabel(btn.Label)
			button.SetFocusable(false)
			buttonsGrid.Attach(button, colId, rowId, 1, 1)

			switch btn.Type {
			case "operator":
				button.StyleContext().AddClass("btn-operator")
			case "reset":
				button.StyleContext().AddClass("btn-reset")
			case "result":
				button.StyleContext().AddClass("btn-result")
			}
		}
	}

	return &buttons
}
