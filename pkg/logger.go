package pkg

import (
	"github.com/rs/zerolog"
	"os"
	"time"
)

type Logger struct {
	logger zerolog.Logger
}

func NewLogger() *Logger {
	output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}

	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	return &Logger{
		logger: zerolog.New(output).With().Timestamp().Logger(),
	}
}

func (l *Logger) Info(msg string) {
	l.logger.Info().Msg(msg)
}

func (l *Logger) Debug(msg string) {
	l.logger.Debug().Msg("custom log " + msg)
}

func (l *Logger) Warn(msg string) {
	l.logger.Warn().Msg("custom log " + msg)
}

func (l *Logger) Error(msg string) {
	l.logger.Error().Msg("custom log " + msg)
}
