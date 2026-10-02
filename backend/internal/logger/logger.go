package logger

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/rs/zerolog"

	"github.com/Din18m/GOnnect/internal/config"
)

func New(config config.LoggerConfig) (zerolog.Logger, func() error, error) {
	level, err := zerolog.ParseLevel(config.Level)
	if err != nil {
		return zerolog.Logger{}, nil, fmt.Errorf("parse logger level: %w", err)
	}

	var out io.Writer = os.Stdout
	closeFn := func() error { return nil }

	if config.Out == "file" {
		f, err := os.OpenFile("logs.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0o644)
		if err != nil {
			return zerolog.Logger{}, nil, fmt.Errorf("open log file: %w", err)
		}
		out = f
		closeFn = f.Close
	}

	if config.Format == "console" {
		out = zerolog.ConsoleWriter{Out: out, TimeFormat: time.RFC3339}
	}

	zerolog.CallerSkipFrameCount = config.SkipFrames

	log := zerolog.New(out).
		Level(level).
		With().
		Caller().
		Timestamp().
		Logger()
	return log, closeFn, nil
}
