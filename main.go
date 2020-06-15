package main

import (
	"github.com/ythosa/gobih/concurrency"
	"github.com/ythosa/gobih/webworkers"
)

func main() {
	go concurrency.StartServer(concurrency.HandleConnEchoServer)
	webworkers.NetCat4EchoServer()
}
