package main

import (
	"time"

	"github.com/ythosa/gobih/logging"
)

var logger, err = logging.GetLogger(logging.LoggerData{Name: "SL", Path: "./log.golang.txt"})

func f1() {
	defer logger.FunctionLogger("f1")()
	time.Sleep(10 * time.Second)
}

func main() {
	if err != nil {
		panic(err)
	}

	f1()
}
