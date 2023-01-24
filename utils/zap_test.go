package utils

import (
	"testing"

	"go.uber.org/zap"
)

func TestName(t *testing.T) {
	url := "Hello"
	logger, _ := zap.NewDevelopment()

	logger.Debug("url", zap.String("url", url))
}
