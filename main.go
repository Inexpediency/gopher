package main

import "github.com/ythosa/gobih/concurrency"

func main() {
	concurrency.RunThreadPool(10000, 500)
}
