package main

import (
	"fmt"
	"github.com/ythosa/gobih/types"
)

func main() {
	if err := types.PageHeaderHTMLChecker("https://github.com"); err != nil {
		fmt.Println(err)
	}
}
