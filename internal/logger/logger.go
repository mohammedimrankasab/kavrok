// Package logger provides structured logging capabilities for Kavrok.
package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// New creates a production-ready structured logger.
func New() (*zap.Logger, error) {
	cfg := zap.NewProductionConfig()

	cfg.Encoding = "json"
	cfg.EncoderConfig.TimeKey = "time"
	cfg.EncoderConfig.LevelKey = "level"
	cfg.EncoderConfig.MessageKey = "message"
	cfg.EncoderConfig.CallerKey = "caller"
	cfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	cfg.EncoderConfig.EncodeLevel = zapcore.LowercaseLevelEncoder

	return cfg.Build(
		zap.AddCaller(),
		zap.AddStacktrace(zap.ErrorLevel),
	)
}
