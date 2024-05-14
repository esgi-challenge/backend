package logger

import (
	"fmt"
	"log/slog"
	"os"
)

type Logger interface {
	InitLogger()
	Debug(msg string, args ...any)
	Info(msg string, args ...any)
	Warn(msg string, args ...any)
	Error(msg string, args ...any)
	Fatal(msg string, args ...any)
	Debugf(format string, args ...any)
	Infof(format string, args ...any)
	Warnf(format string, args ...any)
	Errorf(format string, args ...any)
	Fatalf(format string, args ...any)
}

type logger struct {
	yolog *slog.Logger
}

func NewLogger() *logger {
	return &logger{}
}

func (l *logger) InitLogger() {
	l.yolog = slog.New(slog.NewTextHandler(os.Stdout, nil))
}

func (l *logger) Debug(msg string, args ...any) {
	l.yolog.Debug(msg, args...)
}

func (l *logger) Info(msg string, args ...any) {
	l.yolog.Info(msg, args...)
}

func (l *logger) Warn(msg string, args ...any) {
	l.yolog.Warn(msg, args...)
}

func (l *logger) Error(msg string, args ...any) {
	l.yolog.Error(msg, args...)
}

func (l *logger) Fatal(msg string, args ...any) {
	l.yolog.Error(msg, args...)
	os.Exit(1)
}

func (l *logger) Debugf(format string, args ...any) {
	msg := fmt.Sprintf(format, args...)
	l.yolog.Debug(msg)
}

func (l *logger) Infof(format string, args ...any) {
	msg := fmt.Sprintf(format, args...)
	l.yolog.Info(msg)
}

func (l *logger) Warnf(format string, args ...any) {
	msg := fmt.Sprintf(format, args...)
	l.yolog.Warn(msg)
}

func (l *logger) Errorf(format string, args ...any) {
	msg := fmt.Sprintf(format, args...)
	l.yolog.Error(msg)
}

func (l *logger) Fatalf(format string, args ...any) {
	msg := fmt.Sprintf(format, args...)
	l.yolog.Error(msg)
	os.Exit(1)
}
