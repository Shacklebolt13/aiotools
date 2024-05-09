package widgets

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

func NewGlobalLayout(tabs ...*container.TabItem) fyne.CanvasObject {
	resultTabs := container.NewAppTabs()
	for _, tab := range tabs {
		resultTabs.Append(
			tab,
		)
	}
	return resultTabs
}
