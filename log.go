package logger

import (
	"fmt"
	"log"
	"os"
)

// Defines the log format
const (
	TimeFlag     = log.LstdFlags | log.Lmicroseconds
	FileFlag     = log.Lshortfile
	TimeFileFlag = TimeFlag | FileFlag
)

var l = NewStdLogger(true, true, true, true, TimeFlag)

// Logger A Logger represents an active logging object that generates lines of output to an io.Writer.
type Logger struct {
	logger     *log.Logger
	debug      bool
	trace      bool
	calldepth  int
	infoLabel  string
	errorLabel string
	fatalLabel string
	debugLabel string
	traceLabel string
}

// NewStdLogger creates a logger with output directed to Stderr
func NewStdLogger(debug, trace, colors, pid bool, flag int) *Logger {
	pre := ""
	if pid {
		pre = pidPrefix()
	}

	l := &Logger{
		logger: log.New(os.Stderr, pre, flag),
		debug:  debug,
		trace:  trace,
	}

	if colors {
		setColoredLabelFormats(l)
	} else {
		setPlainLabelFormats(l)
	}

	return l
}

// NewFileLogger creates a logger with output directed to a file
func NewFileLogger(filename string, debug, trace, pid bool, flag int) *Logger {
	fileflags := os.O_WRONLY | os.O_APPEND | os.O_CREATE
	f, err := os.OpenFile(filename, fileflags, 0660)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}

	pre := ""
	if pid {
		pre = pidPrefix()
	}

	l := &Logger{
		logger: log.New(f, pre, flag),
		debug:  debug,
		trace:  trace,
	}

	setPlainLabelFormats(l)
	return l
}

// SetLogger specify one Logger
func SetLogger(logger *Logger) {
	l = logger
}

// generate the pid prefix string
func pidPrefix() string {
	return fmt.Sprintf("[%d] ", os.Getpid())
}

func setPlainLabelFormats(l *Logger) {
	l.infoLabel = "[INF] "
	l.debugLabel = "[DBG] "
	l.errorLabel = "[ERR] "
	l.fatalLabel = "[FTL] "
	l.traceLabel = "[TRC] "
}

func setColoredLabelFormats(l *Logger) {
	colorFormat := "[\x1b[%dm%s\x1b[0m] "
	l.infoLabel = fmt.Sprintf(colorFormat, 32, "INF")
	l.debugLabel = fmt.Sprintf(colorFormat, 36, "DBG")
	l.errorLabel = fmt.Sprintf(colorFormat, 31, "ERR")
	l.fatalLabel = fmt.Sprintf(colorFormat, 31, "FTL")
	l.traceLabel = fmt.Sprintf(colorFormat, 33, "TRC")
}

// SetCallDepth set call depth
func (l *Logger) SetCallDepth(calldepth int) {
	l.calldepth = calldepth
}

// Printf print to the logger
func (l *Logger) Printf(format string, v ...interface{}) {
	l.logger.Output(l.calldepth, fmt.Sprintf(format, v...))
}

// Infof logs a notice statement
func Infof(format string, v ...interface{}) { l.Infof(format, v...) }

// Infof logs a notice statement
func (l *Logger) Infof(format string, v ...interface{}) {
	l.Printf(l.infoLabel+format, v...)
}

// Errorf logs an error statement
func Errorf(format string, v ...interface{}) { l.Errorf(format, v...) }

// Errorf logs an error statement
func (l *Logger) Errorf(format string, v ...interface{}) {
	l.Printf(l.errorLabel+format, v...)
}

// Fatalf logs a fatal error
func Fatalf(format string, v ...interface{}) { l.Fatalf(format, v...) }

// Fatalf logs a fatal error
func (l *Logger) Fatalf(format string, v ...interface{}) {
	l.Fatalf(l.fatalLabel+format, v...)
}

// Debugf logs a debug statement
func Debugf(format string, v ...interface{}) { l.Debugf(format, v...) }

// Debugf logs a debug statement
func (l *Logger) Debugf(format string, v ...interface{}) {
	if l.debug {
		l.Printf(l.debugLabel+format, v...)
	}
}

// Tracef logs a trace statement
func Tracef(format string, v ...interface{}) { l.Tracef(format, v...) }

// Tracef logs a trace statement
func (l *Logger) Tracef(format string, v ...interface{}) {
	if l.trace {
		l.Printf(l.traceLabel+format, v...)
	}
}
