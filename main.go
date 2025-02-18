package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/tfpf/riven/config"
	"github.com/tfpf/riven/logging"
	"log/slog"
	"os"
)

// setUpStructuredLogging ensures that the function name and line number are
// logged along with the usual attributes in JSON to standard output.
func setUpStructuredLogging() {
	handlerOpts := &slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelInfo,
	}
	handler := logging.NewJSONHandler(os.Stdout, handlerOpts)
	logger := slog.New(handler).WithGroup("msg_args")
	slog.SetDefault(logger)
}

func main() {
	setUpStructuredLogging()

	_, _ = config.Read()

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
