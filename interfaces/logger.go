package interfaces

// Logger modules interface, using for dynamic modules
type Logger interface {
	// Create New Logger Object
	New() Logger
	// New Modules init
	// namespace is service name, and version is service version
	Init(namespace string)
	// Set Logging Level, see LogLevel options
	SetLogLevel(level DebugLevel)
	// Get Current Log Level
	GetLogLevel() (level DebugLevel)
	// Set logger output format, default or json
	SetOutputFormat(OutputFormat)
	// Parsing logger message object to string format
	ParsingLog(msg LoggerMessage) (raw string)
	// Same format with fmt.Sprint
	Trace(format interface{}, input ...interface{})
	// Same format with fmt.Sprint
	Debug(format interface{}, input ...interface{})
	// Same format with fmt.Sprint
	Info(format interface{}, input ...interface{})
	// Same format with fmt.Sprint
	Warning(format interface{}, input ...interface{})
	// Same format with fmt.Sprint
	Success(format interface{}, input ...interface{})
	// Same format with fmt.Sprint
	Error(format interface{}, input ...interface{}) Logger
	// Force Exit application
	Quit()
}
