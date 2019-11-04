package logging

import "fmt"

// Logger is app logger
type Logger interface {
	Debug(messag string, args ...interface{})
	Info(message string, args ...interface{})
	Warn(message string, args ...interface{})
	Error(message string, args ...interface{})
}

// LoggerStdOut is log for standard output
type LoggerStdOut struct{}

func (l LoggerStdOut) Debug(message string, args ...interface{}) {
	fmt.Printf("[DEBUG] "+message+"\n", args...)
}

func (l LoggerStdOut) Info(message string, args ...interface{}) {
	fmt.Printf("[INFO] "+message+"\n", args...)
}

func (l LoggerStdOut) Warn(message string, args ...interface{}) {
	fmt.Printf("[WARN] "+message+"\n", args...)
}

func (l LoggerStdOut) Error(message string, args ...interface{}) {
	fmt.Printf("[ERROR] "+message+"\n", args...)
}
