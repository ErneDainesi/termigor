package termigor

import (
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func StartUI(p *os.File) *widget.TextGrid {
	a := app.New()
	w := a.NewWindow("termigor")
	textGrid := widget.NewTextGrid()

	onTypedKey := func(e *fyne.KeyEvent) {
		switch e.Name {
		case fyne.KeyEnter:
			fallthrough
		case fyne.KeyReturn:
			_, _ = p.Write([]byte{'\r'})
		case fyne.KeyBackspace:
			// handle delete
			break
		}
	}

	onTypedRune := func(r rune) {
		_, _ = p.WriteString(string(r))
	}

	w.Canvas().SetOnTypedKey(onTypedKey)
	w.Canvas().SetOnTypedRune(onTypedRune)

	// Create a new container with a wrapped layout
	// set the layout width to 900, height to 325
	w.SetContent(
		fyne.NewContainerWithLayout(
			layout.NewGridWrapLayout(fyne.NewSize(900, 325)),
			textGrid,
		),
	)
	w.ShowAndRun()
    return textGrid
}
