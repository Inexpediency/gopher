package concurrency

import (
	"fmt"
	"log"
	"net"
	"time"
)

// StartServer starts on of servers from
// this package: EchoServer, ClockServer...
func StartServer(fn func (conn net.Conn)) {
	listener, err := net.Listen("tcp", "localhost:8081")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) // For ex.: connection failed
			continue
		}
		go fn(conn)
	}
}


// GoroutineExample shows simple usage goroutines
func GoroutineExample() {
	fmt.Println("Too long counting...")
	go spinner(100 * time.Millisecond)

	const n = 45
	f := fib(n)
	fmt.Printf("\rFibonacci(%d) = %d\n", n, f)
}

func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}

func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)
}
