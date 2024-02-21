package application

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/diamondburned/gotk4/pkg/gdk/v4"
	"github.com/diamondburned/gotk4/pkg/gtk/v4"
)

func (app *Tapplication) HandleConnectButtons() {
	buttons := app.Ui.Buttons

	for _, row := range buttons {
		for _, btn := range row {
			localBtn := btn
			localBtn.Button.ConnectClicked(func() {
				switch localBtn.Type {
				case "reset":
					app.handleReset()
				case "operand":
					app.handleOperand(localBtn.Label)
				case "operator":
					app.handleOperator(localBtn.Label)
				case "percent":
					app.handlePercent()
				case "result":
					app.handleResult()
				}
			})
		}
	}
}

func (app *Tapplication) HandleKeyboard() {
	controller := gtk.NewEventControllerKey()
	controller.ConnectKeyPressed(func(keyval, keycode uint, state gdk.ModifierType) (ok bool) {
		switch keyval {
		case gdk.KEY_KP_0, gdk.KEY_0:
			app.handleOperand("0")
		case gdk.KEY_KP_1, gdk.KEY_1:
			app.handleOperand("1")
		case gdk.KEY_KP_2, gdk.KEY_2:
			app.handleOperand("2")
		case gdk.KEY_KP_3, gdk.KEY_3:
			app.handleOperand("3")
		case gdk.KEY_KP_4, gdk.KEY_4:
			app.handleOperand("4")
		case gdk.KEY_KP_5, gdk.KEY_5:
			app.handleOperand("5")
		case gdk.KEY_KP_6, gdk.KEY_6:
			app.handleOperand("6")
		case gdk.KEY_KP_7, gdk.KEY_7:
			app.handleOperand("7")
		case gdk.KEY_KP_8, gdk.KEY_8:
			app.handleOperand("8")
		case gdk.KEY_KP_9, gdk.KEY_9:
			app.handleOperand("9")
		case gdk.KEY_KP_Decimal, gdk.KEY_period:
			app.handleOperand(".")
		case gdk.KEY_KP_Enter, gdk.KEY_Return:
			app.handleResult()
		case gdk.KEY_KP_Add, gdk.KEY_plus:
			app.handleOperator("+")
		case gdk.KEY_KP_Subtract, gdk.KEY_minus:
			app.handleOperator("-")
		case gdk.KEY_KP_Multiply, gdk.KEY_asterisk:
			app.handleOperator("*")
		case gdk.KEY_KP_Divide, gdk.KEY_slash:
			app.handleOperator("/")
		case gdk.KEY_percent:
			app.handlePercent()
		case gdk.KEY_V:
			app.handleOperator("√")
		case gdk.KEY_Delete, gdk.KEY_BackSpace:
			app.handleReset()
		}
		return ok
	})

	app.Window.AddController(controller)
}

func (app *Tapplication) handleReset() {
	app.Ui.DisplayEntry.SetText("")
	app.Calculator.Operator = ""
	app.Calculator.Operand1 = 0
	app.Calculator.Operand2 = 0
	app.Calculator.AwaitNext = false
}

func (app *Tapplication) handleOperand(operand string) {
	displayEntry := app.Ui.DisplayEntry
	currentText := displayEntry.Text()

	switch operand {
	case ".":
		if len(currentText) > 0 && !strings.Contains(displayEntry.Text(), ".") {
			displayEntry.SetText(currentText + operand)
		}
		break
	case "π":
		if len(currentText) == 0 {
			displayEntry.SetText(fmt.Sprintf("%.9f", math.Pi))
		}
		if app.Calculator.AwaitNext {
			displayEntry.SetText(currentText + fmt.Sprintf("%.9f", math.Pi))
		}
		break
	default:
		displayEntry.SetText(currentText + operand)
	}
}

func (app *Tapplication) handleOperator(value string) {
	displayEntry := app.Ui.DisplayEntry
	currentText := displayEntry.Text()

	if app.Calculator.AwaitNext {
		return
	}

	if len(currentText) > 0 {
		op1, error := strconv.ParseFloat(displayEntry.Text(), 64)

		if error != nil {
			displayEntry.SetText("ERROR")
			return
		}

		app.Calculator.Operand1 = op1
		app.Calculator.Operator = value

		displayEntry.SetText(currentText + value)
		app.Calculator.AwaitNext = true
	}
}

func (app *Tapplication) handlePercent() {
	displayEntry := app.Ui.DisplayEntry
	currentText := displayEntry.Text()

	if strings.ContainsAny(currentText, "+-*/") {
		parts := strings.Split(currentText, app.Calculator.Operator)
		op1, err := strconv.ParseFloat(parts[0], 64)
		op2, err := strconv.ParseFloat(parts[1], 64)
		if err != nil {
			displayEntry.SetText("ERROR")
		}

		result := app.performPercent(op1, op2)
		displayEntry.SetText(fmt.Sprintf("%v", result))

		app.Calculator.Operand1 = result
		app.Calculator.Operand2 = 0
		app.Calculator.AwaitNext = false
	} else {
		op1, err := strconv.ParseFloat(currentText, 64)
		if err != nil {
			displayEntry.SetText("ERROR")
			return
		}

		result := op1 * 0.01
		app.Calculator.Operand1 = result
		app.Calculator.Operand2 = 0
		app.Calculator.AwaitNext = false
		displayEntry.SetText(fmt.Sprintf("%v", result))
	}
}

func (app *Tapplication) handleResult() {
	displayEntry := app.Ui.DisplayEntry
	currentText := displayEntry.Text()
	if app.Calculator.AwaitNext {
		op := strings.Split(currentText, app.Calculator.Operator)[1]
		op2, err := strconv.ParseFloat(op, 64)
		if err != nil {
			displayEntry.SetText("ERROR")
		}
		app.Calculator.Operand2 = op2
		var result float64

		result = app.performCalculation(app.Calculator.Operand1, app.Calculator.Operand2)
		displayEntry.SetText(strconv.FormatFloat(result, 'f', 9, 64))
		app.Calculator.Operand1 = result
		app.Calculator.Operand2 = 0
		app.Calculator.AwaitNext = false
	}
}

func (app *Tapplication) performCalculation(operand1 float64, operand2 float64) float64 {
	switch app.Calculator.Operator {
	case "+":
		return operand1 + operand2
	case "-":
		return operand1 - operand2
	case "*":
		return operand1 * operand2
	case "/":
		return operand1 / operand2
	default:
		return 0
	}
}

func (app *Tapplication) performPercent(operand1 float64, operand2 float64) float64 {
	switch app.Calculator.Operator {
	case "+":
		return (operand1 * operand2 / 100) + operand1
	case "-":
		return operand1 - (operand1 * operand2 / 100)
	case "*":
		return operand1 * operand2 / 100
	case "/":
		return operand1 * 100 / operand2
	default:
		return 0
	}
}
