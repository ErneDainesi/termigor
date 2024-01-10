package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Termigor")

	w.Resize(fyne.NewSize(800, 500))

	tg := widget.NewTextGridFromString(`Hola mundo!
como estan?`)

	cv := w.Canvas()

	cv.SetOnTypedKey(func(ev *fyne.KeyEvent) {
		switch ev.Name {
		case fyne.KeyReturn:
			row := widget.TextGridRow{
				Cells: []widget.TextGridCell{
					{
						Rune:  '0',
						Style: nil,
					},
				},
				Style: nil,
			}
			tg.Rows = append(tg.Rows, row)
			tg.Refresh()
		}
	})

	c := container.NewVBox()

	c.Add(tg)

	w.SetContent(c)
	w.Show()
	a.Run()
}
