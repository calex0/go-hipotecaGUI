package main

import (
	"fmt"
	"math"
	"strconv"

	"github.com/andlabs/ui"
)

func calcularCuota(capital1 float64, interes1 float64, anos1 float64) string {

	var inter float64 = interes1 / 12
	var plazo float64 = anos1 * 12
	var a float64 = 1 + inter/100
	var b float64 = math.Pow(a, -plazo)
	var cuota1 float64 = capital1 * inter / (100 * (1 - b))
	var cuota = fmt.Sprintf("%.2f", cuota1)

	return cuota
}

func main() {
	err := ui.Main(func() {

		window := ui.NewWindow("Cuota hipoteca", 400, 320, false)
		window.SetMargined(true)
		window.OnClosing(func(*ui.Window) bool {
			return true
		})

		capitalBox := ui.NewHorizontalBox()
		capitalBox.SetPadded(true)
		capitalBox.Append(ui.NewLabel("Capital"), false)
		capital := ui.NewEntry()
		capital.SetText("17000")
		capitalBox.Append(capital, true)

		interesBox := ui.NewHorizontalBox()
		interesBox.SetPadded(true)
		interesBox.Append(ui.NewLabel("Interés"), false)
		interes := ui.NewEntry()
		interes.SetText("8.5")
		interesBox.Append(interes, true)

		anosBox := ui.NewHorizontalBox()
		anosBox.SetPadded(true)
		anosBox.Append(ui.NewLabel("Años"), false)
		anos := ui.NewEntry()
		anos.SetText("6")
		anosBox.Append(anos, true)

		labelCuota := ui.NewLabel("")

		buttonBox := ui.NewHorizontalBox()
		buttonBox.SetPadded(true)
		buttonBox.Append(ui.NewLabel(""), true)

		cancel := ui.NewButton("Cancelar")
		cancel.OnClicked(func(*ui.Button) {
			ui.Quit()
		})

		buttonBox.Append(cancel, false)

		calcular := ui.NewButton("Calcular")
		calcular.OnClicked(func(*ui.Button) {

			//cuotaLabel.SetText("")
			var capital1, _ = strconv.ParseFloat(capital.Text(), 64)
			var interes1, _ = strconv.ParseFloat(interes.Text(), 64)
			var anos1, _ = strconv.ParseFloat(anos.Text(), 64)

			capital2 := float64(capital1)
			interes2 := float64(interes1)
			anos2 := float64(anos1)

			var cuota5 = calcularCuota(capital2, interes2, anos2)

			labelCuota.SetText("Cuota mensual: " + cuota5)

		})
		buttonBox.Append(calcular, false)

		layout := ui.NewVerticalBox()
		layout.SetPadded(true)
		layout.Append(capitalBox, false)
		layout.Append(interesBox, false)
		layout.Append(anosBox, false)
		layout.Append(labelCuota, false)

		layout.Append(buttonBox, false)
		window.SetChild(layout)
		window.Show()
	})
	if err != nil {
		panic(err)
	}
}
