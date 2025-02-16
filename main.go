package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/tfpf/riven/config"
	"log/slog"
	"os"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	_, _ = config.NewConfig()
	os.Setenv("FYNE_FONT", `C:\Users\vpaij\AppData\Local\Microsoft\Windows\Fonts\RecMonoCasualNerdFont-Regular.ttf`)
	os.Setenv("FYNE_FONT", `C:\Windows\Fonts\comic.ttf`)
	a := app.New()
	w := a.NewWindow("Hello")

	hello := widget.NewLabel("Hello Fyne!")
	w.SetContent(container.NewVBox(
		hello,
		widget.NewButton("Hi!", func() {
			hello.SetText("huh")
		}),
	))

	w.ShowAndRun()
}
