package echo

import (
	"fmt"
	"os"
	"strings"
)

func wrap(s string) string {
	return fmt.Sprintf("[ %s ]", s)
}

// Start echo
func Start() {
	args := os.Args[1:]

	for i := range args {
		args[i] = wrap(args[i])
	}

	sep := " & "
	fmt.Println("\nArgs are:", strings.Join(args, sep))
}
