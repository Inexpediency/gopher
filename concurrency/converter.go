package concurrency

import (
	"math/rand"
	"sync"
	"time"
)

var used = make(map[int]bool)
var stopped = false
var mu sync.Mutex

func first(min, max int, out chan<- int) {
	for {
		mu.Lock()
		if stopped {
			close(out)
			return
		}
		mu.Unlock()

		out <- rand.Intn(max-min) + min
	}
}

func second(in <-chan int, out chan<- int) {
	for n := range in {
		if used[n] {
			mu.Lock()
			stopped = true
			mu.Unlock()
		} else {
			used[n] = true
			out <- n
		}
	}
	close(out)
}

func third(in <-chan int) {
	sum := 0
	for n := range in {
		sum += n
		print(n, " ")
	}
	println()
	println(sum)
}

func Converter(min, max int) {
	rand.Seed(time.Now().UnixNano())

	f := make(chan int)
	s := make(chan int)

	go first(min, max, f)
	go second(f, s)
	third(s)
}
