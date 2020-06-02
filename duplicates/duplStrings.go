package duplicates

import (
	"bufio"
	"fmt"
	"os"
)

// FindStrings ...
func FindStrings() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)

	isFound := false
	for input.Scan() {
		s := input.Text()
		if s == "" {
			break
		}
		counts[s]++
		if counts[s] > 1 {
			isFound = true
		}
	}

	fmt.Println()

	if !isFound {
		fmt.Println("Duplicates not found.")
		return
	}

	fmt.Println("Duplicates found:")
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
