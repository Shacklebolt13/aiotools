package widgets

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

func NewShortenWidget() *container.TabItem {
	return container.NewTabItem("Shortener", container.NewVBox(
		shortenContent(),
		expandContent(),
	))
}

func shortenContent() fyne.CanvasObject {
	boundLongUrlStringInput := binding.NewString()
	boundLongUrlStringInput.Set("https://www.google.com")
	boundUrlShortOutput := binding.NewString()

	longUrlEntry := widget.NewEntryWithData(boundLongUrlStringInput)
	submitButton := widget.NewButton("Shorten", func() { shortenUrl(boundLongUrlStringInput, boundUrlShortOutput) })
	shortUrlOutputLabel := widget.NewLabelWithData(boundUrlShortOutput)

	layoutTop := container.NewHBox(
		container.NewCenter(
			longUrlEntry,
			submitButton,
		),
	)

	layoutBottom := container.NewCenter(
		shortUrlOutputLabel,
	)

	return container.NewVBox(
		layoutTop,
		layoutBottom,
	)
}

func shortenUrl(longUrl binding.String, shortUrl binding.String) {
	val, err := longUrl.Get()
	if err != nil {
		log.Fatal(err)
	}
	shortUrl.Set(val)
}

func expandContent() fyne.CanvasObject {
	shortInput := binding.NewString()
	shortUrlEntry := widget.NewEntryWithData(shortInput)
	submitButton := widget.NewButton("Expand", func() { expandUrl(shortInput) })
	return container.NewCenter(
		container.NewHBox(
			shortUrlEntry,
			submitButton,
		),
	)
}

func expandUrl(shortUrl binding.String) {
	val, err := shortUrl.Get()
	if err != nil {
		log.Fatal(err)
	}
	shortUrl.Set(val)
}
