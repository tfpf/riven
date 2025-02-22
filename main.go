package main

import (
    "fyne.io/fyne/v2/app"
    "fyne.io/fyne/v2/container"
    "fyne.io/fyne/v2/widget"
    "github.com/tfpf/riven/config"
    "github.com/tfpf/riven/logging"
    "log/slog"
)

func main() {
    slog.SetDefault(logging.NewJSONLogger())

cfg := config.Read()
        cfg.Load()

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
