package logger

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"log/slog"
	"os"

	"github.com/fatih/color"
)

// SetupLogging configures the logging level based on the verbose flag
func SetupLogging(logLvl slog.Level) {
	opts := slog.HandlerOptions{
		Level: logLvl,
	}

	handler := NewPrettyHandler(os.Stdout, opts)
	logger := slog.New(handler)
	slog.SetDefault(logger)
}

// PrettyHandler is a custom slog.Handler that formats log records with colors and prints them to the console
type PrettyHandler struct {
	slog.Handler
	l *log.Logger
}

// Handle formats the log record with colors and prints it to the console
func (h *PrettyHandler) Handle(ctx context.Context, r slog.Record) error {
	level := r.Level.String() + ":"

	switch r.Level {
	case slog.LevelDebug:
		level = color.MagentaString(level)
	case slog.LevelInfo:
		level = color.BlueString(level)
	case slog.LevelWarn:
		level = color.YellowString(level)
	case slog.LevelError:
		level = color.RedString(level)
	}

	timeStr := r.Time.Format("[15:05:05.000]")
	msg := color.CyanString(r.Message)

	fields := make(map[string]any, r.NumAttrs())

	// Convert slog.Record attributes to a map for pretty printing
	r.Attrs(func(a slog.Attr) bool {
		fields[a.Key] = a.Value.Any()
		return true
	})

	if len(fields) == 0 {
		// No fields, just print the message
		h.l.Println(timeStr, level, msg)
		return nil
	}

	b, err := json.Marshal(fields)
	if err != nil {
		return err
	}

	h.l.Println(timeStr, level, msg, color.WhiteString(string(b)))

	return nil
}

func NewPrettyHandler(
	out io.Writer,
	opts slog.HandlerOptions,
) *PrettyHandler {
	return &PrettyHandler{
		Handler: slog.NewJSONHandler(out, &opts),
		l:       log.New(out, "", 0),
	}
}
