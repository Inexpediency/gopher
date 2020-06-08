package main

import (
	"fmt"
	"github.com/ythosa/gobih/types"
)

func main() {
	s := []int{0, 1, 2, 3, 4, 5}
	types.CycleShift(s, 3)
	fmt.Println(s)  // Output: [3 4 5 0 1 2]
}
