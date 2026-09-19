package logging

import "github.com/sudzekai-web-os/core"

type LoggerFactory struct {
	preCategory string
	subCategory string

	writers  []core.ILoggerWriter
	minLevel core.LogLevel
}

func NewLoggerFactory() core.ILoggerFactory {
	return &LoggerFactory{
		writers:  make([]core.ILoggerWriter, 0),
		minLevel: core.LogLevel_NONE,
	}
}

func (lf *LoggerFactory) SetPreCategory(category string) core.ILoggerFactory {
	lf.preCategory = category
	return lf
}

func (lf *LoggerFactory) GetPreCategory() string {
	return lf.preCategory
}

func (lf *LoggerFactory) GetSubCategory() string {
	return lf.subCategory
}

func (lf *LoggerFactory) SetSubCategory(category string) core.ILoggerFactory {
	lf.subCategory = category
	return lf
}

func (lf *LoggerFactory) AddWriter(w core.ILoggerWriter) core.ILoggerFactory {
	lf.writers = append(lf.writers, w)
	return lf
}

func (lf *LoggerFactory) SetMinLevel(level core.LogLevel) {
	lf.minLevel = level
}

func (lf *LoggerFactory) Copy() core.ILoggerFactory {
	writers := make([]core.ILoggerWriter, len(lf.writers))
	copy(writers, lf.writers)

	return &LoggerFactory{
		writers:     writers,
		minLevel:    lf.minLevel,
		preCategory: lf.preCategory,
		subCategory: lf.subCategory,
	}
}

func (lf *LoggerFactory) GetWriters() []core.ILoggerWriter {
	return lf.writers
}

func (lf *LoggerFactory) GetMinLevel() core.LogLevel {
	return lf.minLevel
}

func (lf *LoggerFactory) NewLogger(category string) core.ILogger {
	return &Logger{
		factory:  lf,
		category: category,
	}
}
