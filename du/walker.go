package du

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
)

// walkDir recursively traverses the file tree with the root in dir and sends the size of each found file to file Sizes.
func walkDir(dir string, n *sync.WaitGroup, fileSizes chan<- int64) {
	defer n.Done()
	if cancelled() {
		return
	}

	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			n.Add(1)
			subdir := filepath.Join(dir, entry.Name())
			walkDir(subdir, n, fileSizes)
		} else {
			fileSizes <- entry.Size()
		}
	}
}

// counting semaphore for limiting parallelism
var sema = make(chan struct{}, 20)

// dirents returns directory entries dir
func dirents(dir string) []os.FileInfo {
	select {
	case sema <-struct{}{}: // the capture of the marker
	case <-done: // cancel
		return nil
	}

	// the release of the marker
	defer func(){ <-sema }()

	// getting dir items
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du: %v\n", err)
	}

	return entries
}

// cancelled returns true if counting was cancelled
func cancelled() bool {
	select {
	case <-done:
		return true
	default:
		return false
	}
}
