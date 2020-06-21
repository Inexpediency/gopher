package echo

import (
	"flag"
	"fmt"
	"strings"
)

var (
	n   = flag.Bool("n", false, "Pass new string symbol")
	sep = flag.String("s", " ", "Separator")
)

// StartEchoCLI example: -n -s " Ythosa "  hello world
func StartEchoCLI() {
	flag.Parse()
	s := strings.Join(flag.Args(), *sep)
	if !*n {
		s += "\n"
	}
	fmt.Print(s)
}
