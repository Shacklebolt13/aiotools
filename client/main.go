package main

import (
	"fmt"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

func main() {
	defer tidyUp()
	fmt.Println("Crafting")
	boundData := binding.NewString()

	myApp := app.New()
	myWindow := myApp.NewWindow("Share")
	entryWidget := widget.NewMultiLineEntry()
	entryWidget.Bind(boundData)
	myWindow.SetContent(entryWidget)

	fmt.Println("Starting")
	myWindow.Show()
	myApp.Run()
}

func tidyUp() {
	fmt.Println("Exited")
}
