package concurrency

import (
	"io"
	"log"
	"net"
	"time"
)

func StartClockServer() {
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
		go handleConn(conn)
	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			return // For ex.: user disconnect
		}
		time.Sleep(1 * time.Second)
	}
}
