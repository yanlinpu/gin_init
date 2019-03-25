package logger

import (
	"os"
	"strings"

	log "github.com/sirupsen/logrus"
)

// DefaultLogger log
var (
	DefaultLogger *log.Entry
	logger        *log.Logger
)

func init() {
	level := os.Getenv("LOG_LEVEL")
	env := os.Getenv("GINENV")
	env = strings.ToUpper(env)

	logger = log.New()
	logFormatStr := "text"
	if env == "PRODUCTION" {
		logFormatStr = "json"
		level = "error"
	}

	// 日志格式 json text
	logFormat(logFormatStr)

	// 日志等级
	logLevel(level)

	logger.WithFields(log.Fields{
		"app": "gin_init",
	})
	logger.SetReportCaller(true)
	logger.SetOutput(os.Stdout)
	DefaultLogger = log.NewEntry(logger)
}

func logFormat(format string) {
	format = strings.ToUpper(format)
	switch format {
	case "JSON":
		logger.SetFormatter(&log.JSONFormatter{})
	case "TEXT":
		logger.SetFormatter(&log.TextFormatter{})
	default:
		logger.SetFormatter(&log.TextFormatter{})
	}
}

func logLevel(level string) {
	level = strings.ToUpper(level)
	switch level {
	case "DEBUG":
		logger.SetLevel(log.DebugLevel)
	case "INFO":
		logger.SetLevel(log.InfoLevel)
	case "WARN":
		logger.SetLevel(log.WarnLevel)
	case "ERROR":
		logger.SetLevel(log.ErrorLevel)
	case "FATAL":
		logger.SetLevel(log.FatalLevel)
	default:
		logger.SetLevel(log.InfoLevel)
	}
}

func SetLogLevel(level string) {
	logLevel(level)
}

func SetLogFormat(format string) {
	logFormat(format)
}

// Debug logs a message with debug log level.
func Debug(args ...interface{}) {
	DefaultLogger.Debug(args...)
}

// Debugf logs a formatted message with debug log level.
func Debugf(format string, args ...interface{}) {
	DefaultLogger.Debugf(format, args...)
}

// Info logs a message with info log level.
func Info(args ...interface{}) {
	DefaultLogger.Info(args...)
}

// Infof logs a formatted message with info log level.
func Infof(format string, args ...interface{}) {
	DefaultLogger.Infof(format, args...)
}

// Warn logs a message with warn log level.
func Warn(args ...interface{}) {
	DefaultLogger.Warn(args...)
}

// Warnf logs a formatted message with warn log level.
func Warnf(format string, args ...interface{}) {
	DefaultLogger.Warnf(format, args...)
}

// Error logs a message with error log level.
func Error(args ...interface{}) {
	DefaultLogger.Error(args...)
}

// Errorf logs a formatted message with error log level.
func Errorf(format string, args ...interface{}) {
	DefaultLogger.Errorf(format, args...)
}

// Fatal logs a message with fatal log level.
func Fatal(args ...interface{}) {
	DefaultLogger.Fatal(args...)
}

// Fatalf logs a formatted message with fatal log level.
func Fatalf(format string, args ...interface{}) {
	DefaultLogger.Fatalf(format, args...)
}

// Panic logs a message with panic log level.
func Panic(args ...interface{}) {
	DefaultLogger.Panic(args...)
}

// Panicf logs a formatted message with panic log level.
func Panicf(format string, args ...interface{}) {
	DefaultLogger.Panicf(format, args...)
}
