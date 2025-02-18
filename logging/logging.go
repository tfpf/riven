package logging

import (
	"context"
	"io"
	"log/slog"
)

type JSONHandler struct {
	writer    io.Writer
	addSource bool
	level     slog.Leveler
}

func NewJSONHandler(writer io.Writer, options *slog.HandlerOptions) *JSONHandler {
	handler := &JSONHandler{
		writer:    writer,
		addSource: options.AddSource,
		level:     options.Level,
	}
	if handler.level == nil {
		handler.level = slog.LevelInfo
	}
	return handler
}

// Enabled reports whether handling is done at level.
func (h *JSONHandler) Enabled(_ context.Context, level slog.Level) bool {
	return h.level.Level() <= level
}

// Handle writes record in JSON on a single line.
func (h *JSONHandler) Handle(_ context.Context, record slog.Record) error {
	h.writer.Write([]byte("ok\n"))
	return nil
}

// WithAttrs does nothing.
func (h *JSONHandler) WithAttrs(_ []slog.Attr) slog.Handler {
	return h
}

// WithGroup does nothing.
func (h *JSONHandler) WithGroup(_ string) slog.Handler {
	return h
}
