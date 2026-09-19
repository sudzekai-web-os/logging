package logging

import (
	"fmt"
	"time"

	"github.com/sudzekai-web-os/core"
)

type Logger struct {
	factory  core.ILoggerFactory
	category string
}

func (l *Logger) Log(level core.LogLevel, format string, args ...any) {
	if level < l.factory.GetMinLevel() {
		return
	}

	writers := l.factory.GetWriters()

	for _, writer := range writers {
		writer.Write(core.LogEntry{
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
	l.Log(core.LogLevel_DEBUG, format, args...)
}

func (l *Logger) LogInformation(format string, args ...any) {
	l.Log(core.LogLevel_INFORMATION, format, args...)
}

func (l *Logger) LogWarning(format string, args ...any) {
	l.Log(core.LogLevel_WARNING, format, args...)
}

func (l *Logger) LogError(format string, args ...any) {
	l.Log(core.LogLevel_ERROR, format, args...)
}

func (l *Logger) LogCritical(format string, args ...any) {
	l.Log(core.LogLevel_CRITICAL, format, args...)
}

func (l *Logger) GetCategory() string {
	return l.category
}

func prettifyLogLevel(level core.LogLevel) string {
	switch level {
	case core.LogLevel_NONE:
		return fmt.Sprintf("%-4s", "NONE")
	case core.LogLevel_DEBUG:
		return fmt.Sprintf("\033[30;47m%-4s\033[0m", "DBUG")
	case core.LogLevel_INFORMATION:
		return fmt.Sprintf("\033[34m%-4s\033[0m", "INFO")
	case core.LogLevel_WARNING:
		return fmt.Sprintf("\033[33m%-4s\033[0m", "WARN")
	case core.LogLevel_ERROR:
		return fmt.Sprintf("\033[31m%-4s\033[0m", "ERR")
	case core.LogLevel_CRITICAL:
		return fmt.Sprintf("\033[37;41m%-4s\033[0m", "CRIT")
	default:
		return fmt.Sprintf("\033[37m%-4s\033[0m", "UNKN")
	}
}
