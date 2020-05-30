package webworkers

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

// Fetch ...
func Fetch() {
	for _, url := range os.Args[1:] {
		if !(strings.HasPrefix(url, "https://") || strings.HasPrefix(url, "http://")) {
			url = "https://" + url
		}

		res, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}

		breaker := bufio.NewScanner(os.Stdin)

		fmt.Printf("Request Status: %s\n", res.Status)
		breaker.Scan()
		fmt.Printf("Headers: %s\n", res.Header)
		breaker.Scan()
		fmt.Printf("Cookies: %s\n", res.Cookies())
		breaker.Scan()

		if _, err := io.Copy(os.Stdout, res.Body); err != nil {
			fmt.Fprintf(os.Stderr, "parsing body from req to %s: %v\n", url, err)
		}
	}
}
