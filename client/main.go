package main

import (
	"aiotools/client/internal/widgets"
	"fmt"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
)

var target = os.Getenv("TARGET")

func main() {
	defer tidyUp()
	fmt.Println("Crafting")

	if target == "" {
		target = "localhost:50051"
	}

	client := app.New()
	window := client.NewWindow("AIOTools")
	window.SetMaster()
	window.SetPadded(true)
	window.Resize(fyne.NewSize(1000, 1000))
	tabs := container.NewAppTabs(widgets.NewShortenWidget(), widgets.NewShortenWidget(), widgets.NewShortenWidget())
	// tabs.Resize(fyne.NewSize(800, 600))
	window.SetContent(tabs)

	// grpcClient, err := grpc.NewClient(target)

	// if err != nil {
	// 	fmt.Println("Error creating grpc client")
	// }

	window.ShowAndRun()
}

func tidyUp() {
	fmt.Println("Exited")
}
