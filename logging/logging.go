package logging

import (
	"log"
	"os"
)

// GetLogger returns simple logger
func GetLogger(loggerName, logFilePath string) (*log.Logger, error) {
	f, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}

	ilog := log.New(f, loggerName, log.LstdFlags)
	ilog.SetFlags(log.LstdFlags | log.Lshortfile)

	return ilog, nil
}
