package logging

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/sudzekai-web-os/core"
)

type FileWriter struct {
	openFile func() *os.File
}

func NewFileWriter(fileName string) core.ILoggerWriter {
	return &FileWriter{
		openFile: func() *os.File {
			file, _ := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
			return file
		},
	}
}

func (fw *FileWriter) Write(entry core.LogEntry) {
	if fw.openFile == nil {
		return
	}

	file := fw.openFile()
	if file == nil {
		return
	}

	defer file.Close()

	writeEntry(file, entry)
}

func writeEntry(file *os.File, entry core.LogEntry) {

	fmt.Fprintf(
		file,
		"%s %s %s\n",
		time.Now().Format("15:04:05"),
		prettifyLogLevel(entry.LogLevel),
		entry.Category,
	)

	lines := strings.SplitSeq(entry.Message, "\n")

	for line := range lines {
		fmt.Fprintf(file, "           %s\n", line)
	}
}

func (fw *FileWriter) WriteBatch(entries []core.LogEntry) {
	if fw.openFile == nil {
		return
	}

	file := fw.openFile()
	if file == nil {
		return
	}

	defer file.Close()

	for _, entry := range entries {
		writeEntry(file, entry)
	}
}
