package links

import (
	"fmt"
	"log"
	"os"
)

func BreadthFirst(f func(item string) []string, worklist []string) {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				worklist = append(worklist, f(item)...)
			}
		}
	}
}

func Crawl(url string) []string {
	fmt.Println(url)
	list, err := Extract(url)
	if err != nil {
		log.Print(err)
	}

	return list
}


var tokens = make(chan struct{}, 20)

func CrawlAsync(url string) []string {
	fmt.Println(url)

	tokens <- struct{}{} // the seizure of the marker
	list, err := Extract(url)
	<- tokens            // freeing the marker

	if err != nil {
		log.Print(err)
	}

	return list
}


func Run() {
	BreadthFirst(Crawl, os.Args[1:])
}

func RunAsync() {
	worklist := make(chan []string)
	var n int // Number of waiting to be sent to the list

	// Start with cmd arguments
	n++
	go func() {
		worklist <- os.Args[1:]
	}()

	// Concurrency scan
	seen := make(map[string]bool)
	for ; n > 0; n-- {
		list := <- worklist
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				n++
				go func(link string) {
					worklist <- Crawl(link)
				}(link)
			}
		}
	}
}
