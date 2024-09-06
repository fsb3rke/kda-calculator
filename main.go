package main

import (
	"fmt"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func CalculateKda(kill float32, death float32, assists float32) float32 {
	var m_death float32 = death
	if m_death < 1 {
		m_death = 1
	}

	return (kill + assists) / death
}

func Float32ToString(x float32) string {
	return fmt.Sprintf("%f", x)
}

func StringToFloat32(x string) float32 {
	value, err := strconv.ParseFloat(x, 32)
	if err != nil {
		// do something
	}

	return float32(value)
}

func main() {
	a := app.New()
	w := a.NewWindow("Kda Calculator")

	widget.NewLabel("Kda Calculator")

	kill := widget.NewEntry()
	death := widget.NewEntry()
	assists := widget.NewEntry()

	calculatedLaber := widget.NewLabel("\tYour KDA will be here!")

	w.SetContent(container.NewVBox(
		calculatedLaber,
		widget.NewLabel("Kill"),
		kill,
		widget.NewLabel("Death"),
		death,
		widget.NewLabel("Assists"),
		assists,
		widget.NewLabel(""),
		widget.NewButton("Calculate", func() {
			calculatedLaber.SetText("YOUR KDA IS:\t\t" + Float32ToString(CalculateKda(StringToFloat32(kill.Text), StringToFloat32(death.Text), StringToFloat32(assists.Text))))
		}),
	))

	w.Resize(fyne.NewSize(500, 370))
	w.ShowAndRun()
}
