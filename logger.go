package logger

var defaultLogger = NewStdLogger(true, true, true, true, true)

// SetLogger specify one Logger
func SetLogger(l *Logger) {
	defaultLogger = l
}

// Infof logs a notice statement
func Infof(format string, v ...interface{}) {
	defaultLogger.Infof(format, v...)
}

// Errorf logs an error statement
func Errorf(format string, v ...interface{}) {
	defaultLogger.Errorf(format, v...)
}

// Fatalf logs a fatal error
func Fatalf(format string, v ...interface{}) {
	defaultLogger.Fatalf(format, v...)
}

// Debugf logs a debug statement
func Debugf(format string, v ...interface{}) {
	defaultLogger.Debugf(format, v...)
}

// Tracef logs a trace statement
func Tracef(format string, v ...interface{}) {
	defaultLogger.Tracef(format, v...)
}
