package logger

import (
	"github.com/rs/zerolog"
	"os"
	"time"
)

// Logger is a wrapper around zerolog.Logger that provides structured logging functionality.
type Logger struct {
	logger zerolog.Logger
}

// NewLogger initializes and returns a new instance of Logger with a configured zerolog console writer and timestamp.
func NewLogger() *Logger {
	output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}

	zerolog.SetGlobalLevel(zerolog.GlobalLevel())

	return &Logger{
		logger: zerolog.New(output).With().Timestamp().Logger(),
	}
}

// Info logs an informational message using the embedded zerolog.Logger instance.
func (l *Logger) Info(msg string) {
	l.logger.Info().Msg(msg)
}

// Debug logs a debug-level message with a custom prefix.
func (l *Logger) Debug(msg string) {
	l.logger.Debug().Msg("custom log " + msg)
}

// Warn logs a warning message with a custom prefix.
func (l *Logger) Warn(msg string) {
	l.logger.Warn().Msg("custom log " + msg)
}

// Error logs an error-level message with a custom prefix.
func (l *Logger) Error(msg string) {
	l.logger.Error().Msg("custom log " + msg)
}
