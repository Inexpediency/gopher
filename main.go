package main

import (
	"github.com/ythosa/gobih/logging"
)

func main() {
	logger, err := logging.GetLogger("GLT", "/tmp/golog")
	if err != nil {
		panic(err)
	}

	logger.Println("Hello there!")
	logger.Println("I am here!")
}
