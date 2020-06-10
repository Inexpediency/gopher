package types

import (
	"fmt"
	"sort"
)

func Squares() func() int {
	var x = 0
	return func() int {
		x++
		return x * x
	}
}


func TestSquares() {
	f := Squares()
	fmt.Println(f()) // 1
	fmt.Println(f()) // 4
	fmt.Println(f()) // 9
	fmt.Println(f()) // 16
	fmt.Println(f()) // ...
	fmt.Println(f())
}


func TestTopologySort() {
	var prereqs = map[string][]string{
		"algorithms": {"data structures"},
		"calculus": {"linear algebra"},
		"compilers" : {"data structures", "formal languages", "computer organization"},
		"data structures" : {"discrete math"},
		"databases": {"data structures"},
		"discrete math": {"intro to programming"},
		"formal languages": {"discrete math"},
		"networks": {"operating systems"},
		"programming languages": {"data structures", "computer organization"},
	}

	for i, course := range TopologySort(prereqs) {
		fmt.Printf("%d:\t%s\n", i + 1, course)
	}
}


func TopologySort(m map[string][]string) []string {
	var order []string
	seen := make(map[string]bool)

	var visitAll func(items []string)
	visitAll = func(items []string) {
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				visitAll(m[item])

				order = append(order, item)
			}
		}
	}

	var keys []string
	for key := range m {
		keys = append(keys, key)
	}

	sort.Strings(keys)
	visitAll(keys)

	return order
}

