package main

import (
	"os"

	"github.com/mohammedimrankasab/kavrok/internal/app"
	"github.com/mohammedimrankasab/kavrok/internal/logger"
	"go.uber.org/zap"
)

func main() {
	log, err := logger.New()
	if err != nil {
		log.Error("failed to initialize logger", zap.Error(err))
		os.Exit(1)
	}

	a := app.New(log)

	if err := a.Execute(); err != nil {
		log.Error("application failed", zap.Error(err))
		os.Exit(1)
	}
}
