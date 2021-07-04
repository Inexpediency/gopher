package lexemcounter

import (
	"fmt"
	"go/scanner"
	"go/token"
	"io/ioutil"
)

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
			if token.IsKeyword(lit) {
				localCount[lit] += 1
				count[lit] += 1
			}
		}
		fmt.Println("Found:")
		printCount(localCount)
	}

	printCount(count)
}
