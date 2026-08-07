package app

// App represents the Kavrok application.
type App struct{}

// New creates a new application.
func New() *App {
	return &App{}
}

// Execute starts the application.
func (a *App) Execute() error {
	return nil
}
