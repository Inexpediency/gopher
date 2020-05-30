package duplicates

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// FindStringsFromFileAll ...
func FindStringsFromFileAll() {
	counts := make(map[string]int)
	for _, fileName := os.Args[1:] {
		data, err := ioutil.ReadFile(fileName)
		if err != nil {
			fmt.Fprintf(os.Stderr, "duplicates: %v\n", err)
			continue
		}
		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
		}
	}
	for line, n := range counts {
		fmt.Printf("%d\t%s\n", n, line)
	}
}
