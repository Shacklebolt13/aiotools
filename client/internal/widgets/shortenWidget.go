package widgets

import (
	"aiotools/client/internal/layouts"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func NewShortenWidget() *container.TabItem {
	verticalBox := container.New(
		layouts.NewCustomVBoxLayout(),
		layout.NewSpacer(),
		shortenContent(),
		layout.NewSpacer(),
		expandContent(),
		layout.NewSpacer(),
	)
	verticalBox.Resize(fyne.NewSize(800, 600))
	return container.NewTabItem("Shortener", verticalBox)
}

func shortenContent() fyne.CanvasObject {
	boundLongUrlStringInput := binding.NewString()
	boundUrlShortOutput := binding.NewString()

	longUrlEntry := widget.NewEntryWithData(boundLongUrlStringInput)
	longUrlEntry.SetPlaceHolder("Enter URL to shorten")
	longUrlEntry.Resize(fyne.NewSize(400, 50))

	submitButton := widget.NewButton("Shorten", func() { shortenUrl(boundLongUrlStringInput, boundUrlShortOutput) })
	submitButton.Resize(fyne.NewSize(100, 50))
	shortUrlOutputLabel := widget.NewLabelWithData(boundUrlShortOutput)

	innerTopHBox := container.New(
		layouts.NewCustomHBoxLayout(),
		longUrlEntry,
		layout.NewSpacer(),
		submitButton,
	)
	innerTopHBox.Resize(fyne.NewSize(600, 50))

	layoutTop := container.New(
		layouts.NewCustomHBoxLayout(),
		layout.NewSpacer(),
		innerTopHBox,
		layout.NewSpacer(),
	)
	layoutTop.Resize(fyne.NewSize(800, 50))

	layoutBottom := container.New(
		layouts.NewCustomHBoxLayout(),
		layout.NewSpacer(),
		shortUrlOutputLabel,
		layout.NewSpacer(),
	)
	layoutBottom.Resize(fyne.NewSize(800, 150))

	outerVBox := container.New(
		layouts.NewCustomVBoxLayout(),
		layout.NewSpacer(),
		layoutTop,
		layout.NewSpacer(),
		layoutBottom,
		layout.NewSpacer(),
	)
	outerVBox.Resize(fyne.NewSize(800, 300))

	return outerVBox
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
	shortUrlEntry.SetPlaceHolder("Enter short URL to expand")
	shortUrlEntry.Resize(fyne.NewSize(400, 50))

	submitButton := widget.NewButton("Expand", func() { expandUrl(shortInput) })
	submitButton.Resize(fyne.NewSize(100, 50))

	layoutHBox := container.New(
		layouts.NewCustomHBoxLayout(),
		shortUrlEntry,
		layout.NewSpacer(),
		submitButton,
	)
	layoutHBox.Resize(fyne.NewSize(600, 50))

	innerCenter := container.New(
		layouts.NewCustomHBoxLayout(),
		layout.NewSpacer(),
		layoutHBox,
		layout.NewSpacer(),
	)
	innerCenter.Resize(fyne.NewSize(600, 50))

	outerCenter := container.New(
		layouts.NewCustomVBoxLayout(),
		layout.NewSpacer(),
		innerCenter,
		layout.NewSpacer(),
	)
	outerCenter.Resize(fyne.NewSize(800, 300))

	return outerCenter
}

func expandUrl(shortUrl binding.String) {
	val, err := shortUrl.Get()
	if err != nil {
		log.Fatal(err)
	}
	shortUrl.Set(val)
	log.Default().Println("Expanding URL: ", val)
}
