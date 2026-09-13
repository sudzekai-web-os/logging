package logging

import (
	"fmt"
	"strings"
	"time"

	"github.com/sudzekai/web-os-api/packages/types"
)

type Logger struct {
	factory  *LoggerFactory
	category string
}

func (l *Logger) Log(level types.LogLevel, format string, args ...any) {
	if level < l.factory.GetMinLevel() {
		return
	}

	writer := l.factory.GetWriter()
	if writer == nil {
		return
	}

	fmt.Fprintf(
		writer,
		"%s %s %s\n",
		fmt.Sprintf("\033[90m[%s]\033[0m", time.Now().Format("15:04:05")),
		fmt.Sprintf("%s", prettifyLogLevel(level)),
		fmt.Sprintf("\033[90m%s\033[0m", l.category),
	)

	lines := strings.SplitSeq(fmt.Sprintf(format, args...), "\n")

	for line := range lines {
		fmt.Fprintf(writer, "           %s\n", line)
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
