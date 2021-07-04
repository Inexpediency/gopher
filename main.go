package main

import "github.com/ythosa/gobih/lexemcounter"

func main() {
	lexemcounter.CountVariablesOfLen([]string{"./sudoku/sudoku.go", "./lexemcounter/varlencounter.go"})
}
