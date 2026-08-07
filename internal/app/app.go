// Package app contains the Kavrok application lifecycle and orchestration.
package app

import "go.uber.org/zap"

// App represents the Kavrok application.
type App struct {
	logger *zap.Logger
}

// New creates a new application.
func New(log *zap.Logger) *App {
	return &App{
		logger: log,
	}
}

// Execute starts the application.
func (a *App) Execute() error {
	defer func() {
		_ = a.logger.Sync()
	}()

	a.logger.Info("starting kavrok")

	return nil
}
