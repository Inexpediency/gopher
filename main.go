package main

import (
	"fmt"
	"github.com/ythosa/gobih/links"
)

func main() {
	linksSlice, err := links.Extract("https://github.com/Ythosa/where-is")
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, link := range linksSlice {
		fmt.Println(link)
	}
}
