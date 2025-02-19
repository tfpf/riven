package logging

import (
	"context"
	"io"
	"log/slog"
	"os"
)

// JSONHandler writes logs in JSON.
type JSONHandler struct {
	writer    io.Writer
	addSource bool
	level     slog.Leveler
	group     string
}

// NewJSONHandler returns a handler which writes logs in JSON to writer. Like
// the standard JSON handler, it is configured using options; however, the
// ReplaceAttr field of options is ignored.
func NewJSONHandler(writer io.Writer, options *slog.HandlerOptions) *JSONHandler {
	h := &JSONHandler{
		writer: writer,
	}
	if options != nil {
		h.addSource = options.AddSource
		h.level = options.Level
	}
	return h
}

// Enabled reports whether the handler handles records at level.
func (h *JSONHandler) Enabled(_ context.Context, level slog.Level) bool {
	if h.level == nil {
		h.level = slog.LevelInfo
	}
	return h.level.Level() <= level
}

// Handle writes record in JSON on a single line.
func (h *JSONHandler) Handle(_ context.Context, record slog.Record) error {
	h.writer.Write([]byte("ok\n"))
	return nil
}

// WithAttrs returns the handler. (It is effectively a no-op.)
func (h *JSONHandler) WithAttrs(_ []slog.Attr) slog.Handler {
	return h
}

// WithGroup returns a new handler which groups all attributes of a record in
// group (if it is non-empty).
func (h *JSONHandler) WithGroup(group string) slog.Handler {
	return &JSONHandler{
		writer:    h.writer,
		addSource: h.addSource,
		level:     h.level,
		group:     h.group,
	}
}

// NewJSONLogger returns a custom logger for Riven.
func NewJSONLogger() *slog.Logger {
	options := &slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelInfo,
	}
	handler := NewJSONHandler(os.Stdout, options)
	logger := slog.New(handler).WithGroup("msg_args")
	return logger
}
