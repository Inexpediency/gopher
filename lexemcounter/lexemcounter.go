package lexemcounter

import (
	"fmt"
	"go/scanner"
	"go/token"
	"io/ioutil"
)

func isKeyword(kw string) bool {
	var keywords = []string{
		"break", "default", "func", "interface", "select",
		"case", "defer", "go", "map", "struct",
		"chan", "else", "goto", "package", "switch",
		"const", "fallthrough", "if", "range", "type",
		"continue", "for", "import", "return", "var",
	}

	for _, v := range keywords {
		if kw == v {
			return true
		}
	}

	return false
}

func printCount(c map[string]int) {
	for k, v := range c {
		fmt.Printf("\tlexem: %s\t\tcount: %d\n", k, v)
	}
}

func CountLexem(files []string) {
	count := make(map[string]int)

	for _, file := range files {
		fmt.Println("Processing: ", file)
		data, err := ioutil.ReadFile(file)
		if err != nil {
			fmt.Println(err)
			return
		}
		one := token.NewFileSet()
		files := one.AddFile(file, one.Base(), len(data))

		var s scanner.Scanner
		s.Init(files, data, nil, scanner.ScanComments)

		localCount := make(map[string]int)
		for {
			_, tok, lit := s.Scan()
			if tok == token.EOF {
				break
			}
			if isKeyword(lit) {
				localCount[lit] += 1
				count[lit] += 1
			}
		}
		fmt.Println("Found:")
		printCount(localCount)
	}

	printCount(count)
}
