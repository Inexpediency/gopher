// +build ignore

// The thumbnail command produces thumbnails of JPEG files
// whose names are provided on each line of the standard input.
//
// Run with:
//   - Start `Run()` from main.go
//   - Build program:
//   	$ make
//		$ ./main
//   	$ foo.jpeg
//   	$ ^D
//

package thumbnail

import (
	"bufio"
	"fmt"
	"gopl.io/ch8/thumbnail"
	"log"
	"os"
)

func Run() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		thumb, err := thumbnail.ImageFile(input.Text())
		if err != nil {
			log.Print(err)
			continue
		}
		fmt.Println(thumb)
	}
	if err := input.Err(); err != nil {
		log.Fatal(err)
	}
}