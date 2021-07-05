package concurrency

import (
	"fmt"
	"sync"
	"time"
)

type Client struct {
	id      int
	integer int
}

type Data struct {
	job    Client
	square int
}

var (
	size    = 10
	clients = make(chan Client, size)
	data    = make(chan Data, size)
)

func worker(w *sync.WaitGroup) {
	for c := range clients {
		square := c.integer * c.integer
		output := Data{job: c, square: square}
		data <- output
		time.Sleep(time.Second) // for testing
	}
	w.Done()
}

func makeThreadPool(threadCount int) {
	var w sync.WaitGroup
	for i := 0; i < threadCount; i++ {
		w.Add(1)
		go worker(&w)
	}
	w.Wait()
	close(data)
}

func create(maxN int) {
	for i := 0; i < maxN; i++ {
		c := Client{id: i, integer: i}
		clients <- c
	}
	close(clients)
}

func RunThreadPool(nJobs, nWorkers int) {
	go create(nJobs)
	finished := make(chan interface{})

	go func() {
		for d := range data {
			fmt.Printf("Client ID: %d\tint: ", d.job.id)
			fmt.Printf("%d\tsquare: %d\n", d.job.integer, d.square)
		}
		finished <- true
	}()

	makeThreadPool(nWorkers)
	fmt.Printf(": %v\n", <-finished)
}
