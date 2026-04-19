package logger

import (
	"log/slog"
	"os"
)

type slogLogger struct {
	logger *slog.Logger
}

func NewSlog() Logger {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}))

	return &slogLogger{logger: logger}
}

func (s *slogLogger) Info(msg string, args ...any) {
	s.logger.Info(msg, args...)
}

func (s *slogLogger) Error(msg string, args ...any) {
	s.logger.Error(msg, args...)
}

func (s *slogLogger) Debug(msg string, args ...any) {
	s.logger.Debug(msg, args...)
}

func (s *slogLogger) Warn(msg string, args ...any) {
	s.logger.Warn(msg, args...)
}
