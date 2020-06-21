// Countdown implements the countdown for a rocket launch.

// NOTE: the ticker goroutine never terminates if the launch is aborted.
// This is a "goroutine leak".

package concurrency

import (
	"fmt"
	"os"
	"time"
)

func Countdown() {
	abort := make(chan struct{})
	go func() {
		os.Stdin.Read(make([]byte, 1))
		abort <- struct{}{}
	}()

	fmt.Println("Starting the countdown. Press <Enter> to refuse.")
	tick := time.Tick(1 * time.Second)

	for countdown := 10; countdown > 0; countdown-- {
		fmt.Println(countdown)

		select {
		case <-tick:
			// Do nothing
		}
	}

	launch()
}

func launch() {
	fmt.Println("Lift off!")
}
