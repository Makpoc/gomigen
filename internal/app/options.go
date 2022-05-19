package app

type Option func(*App)

// WithVersion replaces the version, generated based on the module information.
//
// Mainly helpful for tests.
func WithVersion(version string) Option {
	return func(a *App) {
		a.version = version
	}
}

func WithOutputDirectory(path string) Option {
	return func(a *App) {
		a.outputDirectory = path
	}
}

func WithLogger(logger Logger) Option {
	return func(a *App) {
		a.logger = logger
	}
}
