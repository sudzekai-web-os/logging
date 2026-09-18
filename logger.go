package logging

import (
	"fmt"
	"time"

	"github.com/sudzekai-web-os/abstractions"
	"github.com/sudzekai-web-os/types"
)

type Logger struct {
	factory  abstractions.ILoggerFactory
	category string
}

func (l *Logger) Log(level types.LogLevel, format string, args ...any) {
	if level < l.factory.GetMinLevel() {
		return
	}

	writers := l.factory.GetWriters()

	for _, writer := range writers {
		writer.Write(types.LogEntry{
			TimeStamp:   time.Now(),
			LogLevel:    level,
			Message:     fmt.Sprintf(format, args...),
			SubCategory: l.factory.GetSubCategory(),
			Category:    l.category,
			PreCategory: l.factory.GetPreCategory(),
		})
	}
}

func (l *Logger) LogDebug(format string, args ...any) {
	l.Log(types.Debug, format, args...)
}

func (l *Logger) LogInformation(format string, args ...any) {
	l.Log(types.Information, format, args...)
}

func (l *Logger) LogWarning(format string, args ...any) {
	l.Log(types.Warning, format, args...)
}

func (l *Logger) LogError(format string, args ...any) {
	l.Log(types.Error, format, args...)
}

func (l *Logger) LogCritical(format string, args ...any) {
	l.Log(types.Critical, format, args...)
}

func (l *Logger) GetCategory() string {
	return l.category
}

func prettifyLogLevel(level types.LogLevel) string {
	switch level {
	case types.Debug:
		return fmt.Sprintf("\033[30;47m%-4s\033[0m", "DBUG")
	case types.Information:
		return fmt.Sprintf("\033[34m%-4s\033[0m", "INFO")
	case types.Warning:
		return fmt.Sprintf("\033[33m%-4s\033[0m", "WARN")
	case types.Error:
		return fmt.Sprintf("\033[31m%-4s\033[0m", "ERR")
	case types.Critical:
		return fmt.Sprintf("\033[37;41m%-4s\033[0m", "CRIT")
	default:
		return fmt.Sprintf("\033[37m%-4s\033[0m", "UNKN")
	}
}
