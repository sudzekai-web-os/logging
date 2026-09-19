package logging

import (
	"fmt"

	"github.com/sudzekai-web-os/core"
)

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

func levelToString(level core.LogLevel) string {
	switch level {
	case core.LogLevel_NONE:
		return fmt.Sprintf("%-4s", "NONE")
	case core.LogLevel_DEBUG:
		return fmt.Sprintf("%-4s", "DBUG")
	case core.LogLevel_INFORMATION:
		return fmt.Sprintf("%-4s", "INFO")
	case core.LogLevel_WARNING:
		return fmt.Sprintf("%-4s", "WARN")
	case core.LogLevel_ERROR:
		return fmt.Sprintf("%-4s", "ERR")
	case core.LogLevel_CRITICAL:
		return fmt.Sprintf("%-4s", "CRIT")
	default:
		return fmt.Sprintf("%-4s", "UNKN")
	}
}
