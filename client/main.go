package main

import (
	"aiotools/client/internal/widgets"
	"fmt"

	"fyne.io/fyne/v2/app"
)

func main() {
	defer tidyUp()
	fmt.Println("Crafting")
	myApp := app.New()
	widgets.CreateWindow(
		"AIOTools",
		widgets.NewGlobalLayout(
			widgets.NewShortenWidget(),
		),
	)
	myApp.Run()
}

func tidyUp() {
	fmt.Println("Exited")
}
