package logging

import (
	"log"
	"os"
)

// LoggerData struct
type LoggerData struct {
	Name string
	Path string
}

// Logger ...
type Logger struct {
	LoggerData
	lg *log.Logger
}

// GetLogger returns simple logger
func GetLogger(logger LoggerData) (Logger, error) {
	f, err := os.OpenFile(logger.Path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return Logger{}, err
	}

	ilog := log.New(f, logger.Name, log.LstdFlags)
	ilog.SetFlags(log.LstdFlags | log.Lshortfile)

	return Logger{logger, ilog}, nil
}
