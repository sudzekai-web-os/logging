package logging

import (
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/sudzekai-web-os/core"
)

type ConsoleWriter struct {
	w io.Writer
}

func NewConsoleWriter(w io.Writer) core.ILoggerWriter {
	return &ConsoleWriter{
		w: w,
	}
}

func (cw *ConsoleWriter) Write(entry core.LogEntry) {
	if cw.w == nil {
		return
	}

	fmt.Fprintf(
		cw.w,
		"%s %s %s\n",
		fmt.Sprintf("\033[90m[%s]\033[0m", time.Now().Format("15:04:05")),
		fmt.Sprintf("%s", prettifyLogLevel(entry.LogLevel)),
		fmt.Sprintf("\033[90m%s\033[0m", entry.Category),
	)

	lines := strings.SplitSeq(entry.Message, "\n")

	for line := range lines {
		fmt.Fprintf(cw.w, "           %s\n", line)
	}
}

func (cw *ConsoleWriter) WriteBatch(entries []core.LogEntry) {
	for _, entry := range entries {
		cw.Write(entry)
	}
}
