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
