package concurrency

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"time"
)

/*
	simple test by using netcat:
		go concurrency.StartServer(concurrency.HandleConnEchoServer)
		webworkers.NetCat4EchoServer()
*/

func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

func HandleConnEchoServer(c net.Conn) {
	defer c.Close()
	input := bufio.NewScanner(c)
	for input.Scan() {
		go echo(c, input.Text(), 1 * time.Second)
	}
}
