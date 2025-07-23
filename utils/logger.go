package utils

import (
	"fmt"
	"runtime"
	"strings"
	"time"

	"github.com/fatih/color"
)

type LogLevel int

const (
	INFO LogLevel = iota
	WARN
	ERROR
	DEBUG
)

func (l LogLevel) String() string {
	switch l {
	case INFO:
		return "INFO"
	case WARN:
		return "WARN"
	case ERROR:
		return "ERROR"
	case DEBUG:
		return "DEBUG"
	default:
		return "UNKNOWN"
	}
}

func Log(level LogLevel, message string, args ...interface{}) {
	var prefix string
	var logColor func(format string, a ...interface{}) string

	switch level {
	case INFO:
		prefix = "[INFO]"
		logColor = color.New(color.FgGreen).SprintfFunc()
	case WARN:
		prefix = "[WARN]"
		logColor = color.New(color.FgYellow).SprintfFunc()
	case ERROR:
		prefix = "[ERROR]"
		logColor = color.New(color.FgRed).SprintfFunc()
	case DEBUG:
		prefix = "[DEBUG]"
		logColor = color.New(color.FgCyan).SprintfFunc()
	default:
		prefix = "[UNKNOWN]"
		logColor = color.New(color.FgWhite).SprintfFunc()
	}

	pc, file, line, ok := runtime.Caller(2)
	caller := ""
	if ok {
		funcName := runtime.FuncForPC(pc).Name()
		parts := strings.Split(file, "/")
		fileName := parts[len(parts)-1]
		caller = fmt.Sprintf(" (%s:%d %s)", fileName, line, funcName)
	}

	timestamp := time.Now().Format("2006-01-02 15:04:05")
	formattedMessage := fmt.Sprintf(message, args...)
	fmt.Printf("%s %s %s%s\n", timestamp, logColor("%s", prefix), formattedMessage, color.New(color.FgHiBlack).SprintfFunc()("%s", caller))
}
