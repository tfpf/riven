package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/tfpf/riven/config"
	"log/slog"
	"os"
)

// setUpStructuredLogging ensures that the function name and line number are
// logged along with the usual attributes in JSON to standard output.
func setUpStructuredLogging() {
	handlerOpts := &slog.HandlerOptions{
		AddSource: true,
		ReplaceAttr: func(_ []string, attr slog.Attr) slog.Attr {
			if attr.Key == slog.SourceKey {
				if source, ok := attr.Value.Any().(*slog.Source); ok {
					// The function name tells us which file it is from, so
					// don't log the file.
					source.File = ""
				}
			}
			return attr
		},
	}
	handler := slog.NewJSONHandler(os.Stdout, handlerOpts)
	logger := slog.New(handler)
	slog.SetDefault(logger)
}

func main() {
	setUpStructuredLogging()

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
