package concurrency

import "fmt"

func SquaresPipeline() {
	naturals := make(chan int)
	squares := make(chan int)

	go func() {
		for i := 0; ; i++ {
			naturals <- i
		}
	}()

	go func() {
		for {
			x := <- naturals
			squares <- x * x
		}
	}()

	for {
		x := <- squares
		fmt.Printf("%d  ", x)
	}
}


func SquaresPipelineWithClosinChannels() {
	naturals := make(chan int)
	squares := make(chan int)

	// Generate numbers
	go func() {
		for i := 0; i < 100; i++ {
			naturals <- i
		}
		close(naturals)
	}()

	// Square numbers
	go func() {
		for x := range naturals {
			squares <- x * x
		}
		close(squares)
	}()

	// Output numbers (in main goroutine)
	for x := range squares {
		fmt.Printf("%d  ", x)
	}
}


func counter(out chan<- int) {
	for x := 0; x < 100; x++ {
		out <- x
	}
	close(out)
}

func squarer(out chan<-int, in <-chan int) {
	for x := range in {
		out <- x * x
	}
	close(out)
}

func printer(in <-chan int) {
	for x := range in {
		fmt.Printf("%d  ", x)
	}
}

func SquaresPipelineUnidirectional() {
	naturals := make(chan int)
	squares := make(chan int)

	go counter(naturals)
	go squarer(squares, naturals)
	printer(squares)
}
