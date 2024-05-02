package logger_test

import (
	"testing"

	"github.com/esgi-challenge/backend/pkg/logger"
)

func TestLogger(t *testing.T) {
	logger := logger.NewLogger()
	logger.InitLogger()

	logger.Debug("This is a debug message")
	logger.Info("This is an info message")
	logger.Warn("This is a warning message")
	logger.Error("This is an error message")
}
