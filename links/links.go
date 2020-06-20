package links

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"os"
)

// Extract sends GET HTTP request to the URL, completes
// syntax analise and returns all links in HTML document
func Extract(url string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("getting %s by HTML: %v", url, err)
	}

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("analise %s by HTML: %v", url, err)
	}

	var links []string
	visitNode := func (n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key != "href" {
					continue // Ignore other props
				}

				link, err := resp.Request.URL.Parse(a.Val)
				if err != nil {
					continue // Ignore not correct URLs
				}

				links = append(links, link.String())
			}
		}
	}

	ForEachNode(doc, visitNode, nil)
	return links, nil
}


func TestForEachNode() {
	f, err := os.Open("./links/templatehtmlprint.html")
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
