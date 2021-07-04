package main

import (
	"os"

	"github.com/ythosa/gobih/unix"
)

func main() {
	unix.TraceSysCalls(os.Args)
}
