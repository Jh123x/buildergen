package common

// Logger defines a minimal interface for logging within the application.
type Logger interface {
	Infof(format string, args ...any)
	Errorf(format string, args ...any)
	Debugf(format string, args ...any)
}
