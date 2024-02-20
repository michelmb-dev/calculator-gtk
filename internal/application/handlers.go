package application

import (
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
	"github.com/diamondburned/gotk4/pkg/gtk/v4"
)

func (c *Tcalculator) HandleConnectButtons() {
	buttons := c.Ui.Buttons

	for _, row := range buttons {
		for _, btn := range row {
			localBtn := btn
			localBtn.Button.ConnectClicked(func() {
				switch localBtn.Type {
				case "reset":
					c.handleReset()
				case "operand":
					c.handleOperand(localBtn.Label)
				case "operator":
					c.handleOperator(localBtn.Label)
				case "result":
					c.handleResult()
				}
			})
		}
	}
}

func (c *Tcalculator) HandleKeyboard() {
	controller := gtk.NewEventControllerKey()
	controller.ConnectKeyPressed(func(keyval, keycode uint, state gdk.ModifierType) (ok bool) {
		switch keyval {
		case gdk.KEY_KP_0, gdk.KEY_0:
			c.handleOperand("0")
		case gdk.KEY_KP_1, gdk.KEY_1:
			c.handleOperand("1")
		case gdk.KEY_KP_2, gdk.KEY_2:
			c.handleOperand("2")
		case gdk.KEY_KP_3, gdk.KEY_3:
			c.handleOperand("3")
		case gdk.KEY_KP_4, gdk.KEY_4:
			c.handleOperand("4")
		case gdk.KEY_KP_5, gdk.KEY_5:
			c.handleOperand("5")
		case gdk.KEY_KP_6, gdk.KEY_6:
			c.handleOperand("6")
		case gdk.KEY_KP_7, gdk.KEY_7:
			c.handleOperand("7")
		case gdk.KEY_KP_8, gdk.KEY_8:
			c.handleOperand("8")
		case gdk.KEY_KP_9, gdk.KEY_9:
			c.handleOperand("9")
		case gdk.KEY_KP_Decimal, gdk.KEY_period:
			c.handleOperand(".")
		case gdk.KEY_KP_Enter, gdk.KEY_Return:
			c.handleResult()
		case gdk.KEY_KP_Add, gdk.KEY_plus:
			c.handleOperator("+")
		case gdk.KEY_KP_Subtract, gdk.KEY_minus:
			c.handleOperator("-")
		case gdk.KEY_KP_Multiply, gdk.KEY_asterisk:
			c.handleOperator("*")
		case gdk.KEY_KP_Divide, gdk.KEY_slash:
			c.handleOperator("/")
		case gdk.KEY_percent:
			c.handleOperator("%")
		case gdk.KEY_V:
			c.handleOperator("√")
		case gdk.KEY_Delete, gdk.KEY_BackSpace:
			c.handleReset()
		}
		return ok
	})

	c.Window.AddController(controller)
}

func (c *Tcalculator) handleReset() {
	c.Ui.DisplayEntry.SetText("")
}

func (c *Tcalculator) handleOperand(value string) {
	displayEntry := c.Ui.DisplayEntry
	currentText := displayEntry.Text()

	displayEntry.SetText(currentText + value)
}

func (c *Tcalculator) handleOperator(value string) {
	displayEntry := c.Ui.DisplayEntry
	currentText := displayEntry.Text()

	displayEntry.SetText(currentText + value)
}

// handleResult  TODO: evaluate expression for calculs
func (c *Tcalculator) handleResult() {
}
