package types

import (
	"fmt"
	"golang.org/x/net/html"
	"os"
)

func TestForEachNode() {
	f, err := os.Open("./types/templatehtmlprint.html")
	if err != nil {
		fmt.Println(err)
		return
	}

	h, err := html.Parse(f)
	fmt.Println(h)
	if err != nil {
		fmt.Println(err)
		return
	}

	ForEachNode(h, StartElement, EndElement)
}


func ForEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		ForEachNode(c, pre, post)
	}
	if post != nil {
		post(n)
	}
}

var depth = 0

func StartElement(n *html.Node) {
	if n.Type == html.ElementNode {
		fmt.Printf("%*s<%s>\n", depth*4, "", n.Data)
		depth++
	}
}

func EndElement(n *html.Node) {
	if n.Type == html.ElementNode {
		depth--
		fmt.Printf("%*s</%s>\n", depth*4, "", n.Data)
	}
}
