package interfaces

type Logger interface {
	// Create New Logger Object
	New() Logger

	// New Modules init
	// namespace is service name
	Init() Logger

	// Set logger output format, text or json,
	// Use `text` for Text Formatter,
	// Use `json` for JSON Formater
	SetFormat(format string) Logger

	// Set Logging File
	SetLevel(level string) Logger

	// Flag for whether to log caller info (off by default)
	SetReportCaller(reportCaller bool)

	// InfoLevel level. General operational entries about what's going on inside the application,
	// And auto parsing message struct to json
	Info(msg interface{}, opts ...interface{}) Logger

	// InfoLevel level. General operational entries about what's going on inside the application,
	Infof(msg string, opts ...interface{}) Logger

	// WarnLevel level. Non-critical entries that deserve eyes.
	// And auto parsing message struct to json
	Warn(msg interface{}, opts ...interface{}) Logger

	// WarnLevel level. Non-critical entries that deserve eyes.
	Warnf(msg string, opts ...interface{}) Logger

	// DebugLevel level. Usually only enabled when debugging. Very verbose logging.
	// And auto parsing message struct to json
	Debug(msg interface{}, opts ...interface{}) Logger

	// DebugLevel level. Usually only enabled when debugging. Very verbose logging.
	Debugf(msg string, opts ...interface{}) Logger

	// ErrorLevel level. Logs. Used for errors that should definitely be noted.
	// Commonly used for hooks to send errors to an error tracking service.
	// And auto parsing message struct to json
	Error(msg interface{}, opts ...interface{}) Logger

	// ErrorLevel level. Logs. Used for errors that should definitely be noted.
	// Commonly used for hooks to send errors to an error tracking service.
	Errorf(msg string, opts ...interface{}) Logger

	// FatalLevel level. Logs and then calls `logger.Exit(1)`. It will exit even if the
	// logging level is set to Panic.
	// And auto parsing message struct to json
	Fatal(msg interface{}, opts ...interface{}) Logger

	// FatalLevel level. Logs and then calls `logger.Exit(1)`. It will exit even if the
	// logging level is set to Panic.
	Fatalf(msg string, opts ...interface{}) Logger
}
