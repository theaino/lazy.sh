package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"
)

func init() {
	slog.SetDefault(slog.New(NewLogHandler(os.Stderr, slog.LevelInfo)))
}

type LogHandler struct {
	writer *os.File
	level  slog.Level
}

func NewLogHandler(w *os.File, level slog.Level) *LogHandler {
	return &LogHandler{
		writer: w,
		level:  level,
	}
}

func (h *LogHandler) Enabled(_ context.Context, level slog.Level) bool {
	return level >= h.level
}

func (h *LogHandler) Handle(_ context.Context, r slog.Record) error {
	level := "[" + r.Level.String() + "] "
	if r.Level == slog.LevelInfo {
		level = ""
	}
	msg := r.Message

	fmt.Fprintf(h.writer, "lazysh: %s%s\n", level, msg)
	return nil
}

func (h *LogHandler) WithAttrs(_ []slog.Attr) slog.Handler {
	return h // ignored
}

func (h *LogHandler) WithGroup(_ string) slog.Handler {
	return h // ignored
}
