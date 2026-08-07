package main

import (
	"log"
	"os"

	"go.uber.org/zap"

	"github.com/mohammedimrankasab/kavrok/internal/app"
	"github.com/mohammedimrankasab/kavrok/internal/logger"
)

// App represents the Kavrok application.
type App struct {
	logger *zap.Logger
}

// New creates a new application.
func New() *App {
	log, err := logger.New()
	if err != nil {
		panic(err)
	}

	return &App{
		logger: log,
	}
}

// Execute starts the application.
func (a *App) Execute() error {
	a.logger.Info("starting kavrok")

	defer func() {
		_ = a.logger.Sync()
	}()

	return nil
}

func main() {
	a := app.New()

	if err := a.Execute(); err != nil {
		log.Printf("application failed: %v", err)
		os.Exit(1)
	}
}
