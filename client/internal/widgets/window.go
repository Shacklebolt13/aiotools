package widgets

import "fyne.io/fyne/v2"

func NewWindow(title string, content fyne.CanvasObject) fyne.Window {
	window := fyne.CurrentApp().NewWindow(title)
	window.SetContent(content)
	return window
}
