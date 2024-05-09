package widgets

import "fyne.io/fyne/v2"

func CreateWindow(title string, content fyne.CanvasObject) {
	window := fyne.CurrentApp().NewWindow(title)
	window.SetContent(content)
	window.Show()
}
