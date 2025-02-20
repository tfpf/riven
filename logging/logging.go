package logging

import (
	"context"
	"encoding/json"
	"io"
	"log/slog"
	"os"
	"runtime"
	"time"
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
	details := map[string]any{
		"level": record.Level,
		"msg":   record.Message,
		"time":  record.Time.Format(time.RFC3339),
	}
	if h.addSource {
		frame, _ := runtime.CallersFrames([]uintptr{record.PC}).Next()
		details["source"] = map[string]any{
			"function": frame.Function,
			"line":     frame.Line,
		}
	}
	detailsGroup := details
	if h.group != "" {
		detailsGroup = make(map[string]any)
		details[h.group] = detailsGroup
	}
	record.Attrs(func(attr slog.Attr) bool {
		detailsGroup[attr.Key] = attr.Value.Any()
		return true
	})
	detailsBytes, err := json.Marshal(details)
	if err != nil {
		return err
	}
	detailsBytes = append(detailsBytes, '\n')
	_, err = h.writer.Write(detailsBytes)
	return err
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
		group:     group,
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
