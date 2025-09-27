package logger

import (
	"context"
	"log/slog"
	"os"
	"strings"
)

type Logger struct {
	logLevel slog.Level
	log      *slog.Logger
}

func NewLogger(logLevel string) *Logger {

	l := slog.New(slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{Level: StringToLogLevel(logLevel)}))
	return &Logger{
		logLevel: StringToLogLevel(logLevel),
		log:      l,
	}
}

func (l Logger) GetLogger() *slog.Logger {
	return l.log
}

func (l Logger) GetLogLevel() slog.Level {
	return l.logLevel
}

func StringToLogLevel(i string) slog.Level {
	level := strings.ToLower(i)
	switch level {
	case "debug":
		return slog.LevelDebug
	case "info":
		return slog.LevelInfo
	case "warn":
		return slog.LevelWarn
	case "error":
		return slog.LevelError
	default:
		return slog.LevelDebug
	}
}

func (l *Logger) Debug(ctx context.Context, msg string, args ...any) {
	l.log.DebugContext(ctx, msg, args...)
}

func (l *Logger) Info(ctx context.Context, msg string, args ...any) {
	l.log.InfoContext(ctx, msg, args...)
}

func (l *Logger) Warn(ctx context.Context, msg string, args ...any) {
	l.log.WarnContext(ctx, msg, args...)
}

func (l *Logger) Error(ctx context.Context, msg string, args ...any) {
	l.log.ErrorContext(ctx, msg, args...)
}

// Print logs a message at the Info level.
func (l *Logger) Print(msg string, args ...any) {
	l.log.Info(msg, args...)
}

func (l *Logger) PrintErrorMessage(msg string, args ...any) {
	l.log.Error(msg, args...)
}

func (l *Logger) PrintError(err error) {
	l.log.Error(err.Error())
}
