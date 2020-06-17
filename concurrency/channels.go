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
