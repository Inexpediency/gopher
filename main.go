package main

import (
	"github.com/ythosa/gobih/chat"
	"github.com/ythosa/gobih/webworkers"
)

func main() {
	go chat.Start()
	webworkers.NetCat4EchoServer()
}
