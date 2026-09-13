package logging

import (
	"fmt"
	"io"
	"strings"
	"sync"

	"github.com/sudzekai-web-os/abstractions"
	"github.com/sudzekai-web-os/types"
)

type LoggerFactory struct {
	mu sync.RWMutex

	writer   io.Writer
	minLevel types.LogLevel
}

func NewLoggerFactory(writer io.Writer) abstractions.ILoggerFactory {
	return &LoggerFactory{
		writer:   writer,
		minLevel: types.Information,
	}
}

func (f *LoggerFactory) SetWriter(writer io.Writer) {
	f.mu.Lock()
	defer f.mu.Unlock()

	f.writer = writer
}

func (f *LoggerFactory) SetMinLevel(level types.LogLevel) {
	f.mu.Lock()
	defer f.mu.Unlock()

	f.minLevel = level
}

func (f *LoggerFactory) SetMinLevelStr(level string) error {
	f.mu.Lock()
	defer f.mu.Unlock()

	parsed, err := parseLogLevel(level)
	if err != nil {
		return err
	}

	f.minLevel = parsed
	return nil
}

func (f *LoggerFactory) GetWriter() io.Writer {
	f.mu.RLock()
	defer f.mu.RUnlock()

	return f.writer
}

func (f *LoggerFactory) GetMinLevel() types.LogLevel {
	f.mu.RLock()
	defer f.mu.RUnlock()

	return f.minLevel
}

func (f *LoggerFactory) NewLogger(category string) abstractions.ILogger {
	return &Logger{
		factory:  f,
		category: category,
	}
}

func parseLogLevel(level string) (types.LogLevel, error) {
	switch strings.ToLower(strings.TrimSpace(level)) {
	case "debug", "dbug":
		return types.Debug, nil
	case "information", "info":
		return types.Information, nil
	case "warning", "warn":
		return types.Warning, nil
	case "error", "err":
		return types.Error, nil
	case "critical", "crit":
		return types.Critical, nil
	default:
		return 0, fmt.Errorf("неизвестный уровень логирования: %q", level)
	}
}
