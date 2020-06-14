package main

import (
	"fmt"
	"github.com/ythosa/gobih/interfaces"
)

func main() {
	if err := interfaces.SortTasks("artist", false); err != nil {
		fmt.Println(err)
	}
}
