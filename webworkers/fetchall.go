package webworkers

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

// FetchAll ...
func FetchAll() {
	start := time.Now()
	ch := make(chan string)

	urls := os.Args[1:]

	for _, url := range urls {
		go fetching(setPrefix(url), ch)
	}

	for range urls {
		fmt.Println(<-ch)
	}

	fmt.Printf("Total elapsed: %.3fs\n", time.Since(start).Seconds())
}

func fetching(url string, ch chan string) {
	start := time.Now()
	res, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
	}

	nbytes, err := io.Copy(ioutil.Discard, res.Body)
	res.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}

	secs := time.Since(start).Seconds()

	ch <- fmt.Sprintf("Time: %.3fs | Size: %7d | URL: %s", secs, nbytes, url)
}
