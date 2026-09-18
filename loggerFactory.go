package logging

import (
	"github.com/sudzekai-web-os/abstractions"
	"github.com/sudzekai-web-os/types"
)

type LoggerFactory struct {
	preCategory string
	subCategory string

	writers  []abstractions.ILoggerWriter
	minLevel types.LogLevel
}

func NewLoggerFactory() abstractions.ILoggerFactory {
	return &LoggerFactory{
		writers:  make([]abstractions.ILoggerWriter, 0),
		minLevel: types.None,
	}
}

func (lf *LoggerFactory) SetPreCategory(category string) abstractions.ILoggerFactory {
	lf.preCategory = category
	return lf
}

func (lf *LoggerFactory) GetPreCategory() string {
	return lf.preCategory
}

func (lf *LoggerFactory) GetSubCategory() string {
	return lf.subCategory
}

func (lf *LoggerFactory) SetSubCategory(category string) abstractions.ILoggerFactory {
	lf.subCategory = category
	return lf
}

func (lf *LoggerFactory) AddWriter(w abstractions.ILoggerWriter) abstractions.ILoggerFactory {
	lf.writers = append(lf.writers, w)
	return lf
}

func (lf *LoggerFactory) SetMinLevel(level types.LogLevel) {
	lf.minLevel = level
}

func (lf *LoggerFactory) Copy() abstractions.ILoggerFactory {
	writers := make([]abstractions.ILoggerWriter, len(lf.writers))
	copy(writers, lf.writers)

	return &LoggerFactory{
		writers:     writers,
		minLevel:    lf.minLevel,
		preCategory: lf.preCategory,
		subCategory: lf.subCategory,
	}
}

func (lf *LoggerFactory) GetWriters() []abstractions.ILoggerWriter {
	return lf.writers
}

func (lf *LoggerFactory) GetMinLevel() types.LogLevel {
	return lf.minLevel
}

func (lf *LoggerFactory) NewLogger(category string) abstractions.ILogger {
	return &Logger{
		factory:  lf,
		category: category,
	}
}
