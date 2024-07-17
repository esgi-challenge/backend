package logger

import (
	"bytes"
	"context"
	"os"
	"os/exec"
	"testing"

	"github.com/stretchr/testify/assert"
	"log/slog"
)

type TestHandler struct {
	buffer bytes.Buffer
}

func (h *TestHandler) Handle(ctx context.Context, r slog.Record) error {
	h.buffer.WriteString(r.Message)
	return nil
}

func (h *TestHandler) Enabled(ctx context.Context, level slog.Level) bool {
	return true
}

func (h *TestHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return h
}

func (h *TestHandler) WithGroup(name string) slog.Handler {
	return h
}

func setupTestLogger() (*logger, *TestHandler) {
	l := NewLogger()
	handler := &TestHandler{}
	l.yolog = slog.New(handler)
	return l, handler
}

func TestLogger_Debug(t *testing.T) {
	l, handler := setupTestLogger()
	l.Debug("Debug message")
	assert.Contains(t, handler.buffer.String(), "Debug message")
}

func TestLogger_Info(t *testing.T) {
	l, handler := setupTestLogger()
	l.Info("Info message")
	assert.Contains(t, handler.buffer.String(), "Info message")
}

func TestLogger_Warn(t *testing.T) {
	l, handler := setupTestLogger()
	l.Warn("Warn message")
	assert.Contains(t, handler.buffer.String(), "Warn message")
}

func TestLogger_Error(t *testing.T) {
	l, handler := setupTestLogger()
	l.Error("Error message")
	assert.Contains(t, handler.buffer.String(), "Error message")
}

func TestLogger_Fatal(t *testing.T) {
	if os.Getenv("BE_CRASHER") == "1" {
		l, _ := setupTestLogger()
		l.Fatal("Fatal message")
		return
	}
	cmd := exec.Command(os.Args[0], "-test.run=TestLogger_Fatal")
	cmd.Env = append(os.Environ(), "BE_CRASHER=1")
	err := cmd.Run()
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "exit status 1")
}

func TestLogger_Debugf(t *testing.T) {
	l, handler := setupTestLogger()
	l.Debugf("Debug %s", "formatted message")
	assert.Contains(t, handler.buffer.String(), "Debug formatted message")
}

func TestLogger_Infof(t *testing.T) {
	l, handler := setupTestLogger()
	l.Infof("Info %s", "formatted message")
	assert.Contains(t, handler.buffer.String(), "Info formatted message")
}

func TestLogger_Warnf(t *testing.T) {
	l, handler := setupTestLogger()
	l.Warnf("Warn %s", "formatted message")
	assert.Contains(t, handler.buffer.String(), "Warn formatted message")
}

func TestLogger_Errorf(t *testing.T) {
	l, handler := setupTestLogger()
	l.Errorf("Error %s", "formatted message")
	assert.Contains(t, handler.buffer.String(), "Error formatted message")
}

func TestLogger_Fatalf(t *testing.T) {
	if os.Getenv("BE_CRASHER") == "1" {
		l, _ := setupTestLogger()
		l.Fatalf("Fatal %s", "formatted message")
		return
	}
	cmd := exec.Command(os.Args[0], "-test.run=TestLogger_Fatalf")
	cmd.Env = append(os.Environ(), "BE_CRASHER=1")
	err := cmd.Run()
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "exit status 1")
}
