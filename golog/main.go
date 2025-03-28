package golog

import (
    "log"
    "os"
)

// Define log levels
const (
    DEBUG = iota
    INFO
    WARN
    ERROR
)

// Current log level
var currentLogLevel = INFO

// Custom logger instances
var (
    debugLogger = log.New(os.Stdout, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile)
    infoLogger  = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime)
    warnLogger  = log.New(os.Stdout, "WARN: ", log.Ldate|log.Ltime)
    errorLogger = log.New(os.Stdout, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
)

// SetLogLevel sets the current log level
func SetLogLevel(level int) {
    currentLogLevel = level
}

// Debug logs a message at DEBUG level
func Debug(v ...any) {
    if currentLogLevel <= DEBUG {
        debugLogger.Println(v...)
    }
}

// Debugf logs a formatted message at DEBUG level
func Debugf(format string, v ...any) {
    if currentLogLevel <= DEBUG {
        debugLogger.Printf(format, v...)
    }
}

// Info logs a message at INFO level
func Info(v ...any) {
    if currentLogLevel <= INFO {
        infoLogger.Println(v...)
    }
}

// Infof logs a formatted message at INFO level
func Infof(format string, v ...any) {
    if currentLogLevel <= INFO {
        infoLogger.Printf(format, v...)
    }
}

// Warn logs a message at WARN level
func Warn(v ...any) {
    if currentLogLevel <= WARN {
        warnLogger.Println(v...)
    }
}

// Warnf logs a formatted message at WARN level
func Warnf(format string, v ...any) {
    if currentLogLevel <= WARN {
        warnLogger.Printf(format, v...)
    }
}

// Error logs a message at ERROR level
func Error(v ...any) {
    if currentLogLevel <= ERROR {
        errorLogger.Println(v...)
    }
}

// Errorf logs a formatted message at ERROR level
func Errorf(format string, v ...any) {
    if currentLogLevel <= ERROR {
        errorLogger.Printf(format, v...)
    }
}
