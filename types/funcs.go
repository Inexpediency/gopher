package types

import (
	"fmt"
	"github.com/ythosa/gobih/links"
	"golang.org/x/net/html"
	"net/http"
	"sort"
	"strings"
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

func Sum(values ...int) int {
	result := 0
	for v := range values {
		result += v
	}
	return result
}

func TestSum() {
	Sum(1, 2, 3)

	values := make([]int, 0)
	values = append(values, 1, 2, 3)
	Sum(values...)
}

func PageHeaderHTMLChecker(url string) error {
	res, err := http.Get(url)
	if err != nil {
		return err
	}

	defer res.Body.Close()
	ct := res.Header.Get("Content-Type")
	if ct != "text/html" && !strings.HasPrefix(ct, "text/html") {
		return fmt.Errorf("%s has type %s, but not text/html", url, ct)
	}

	doc, err := html.Parse(res.Body)
	if err != nil {
		return fmt.Errorf("analis %s how HTML: %v", url, err)
	}

	visitNode := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "title" && n.FirstChild != nil {
			fmt.Println(n.FirstChild.Data)
		}
	}

	links.ForEachNode(doc, visitNode, nil)

	return nil
}
