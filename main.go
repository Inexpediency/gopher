package main

import (
	"github.com/ythosa/gobih/web"
)

func main() {
	web.StartKeyValueStorageServer(":8080")
}
