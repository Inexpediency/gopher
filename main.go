package main

import "github.com/ythosa/gobih/bases/webworkers"

func main() {
	go webworkers.StartServer()
	webworkers.Fetch()
}
