package concurrency

import (
	"io"
	"net"
	"time"
)

/*
	simple test by using netcat:
		go concurrency.StartServer(concurrency.HandleConnClockServer)
		webworkers.NetCat()
*/

// HandleConnClockServer starts clock server
func HandleConnClockServer(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			return // For ex.: user disconnect
		}
		time.Sleep(1 * time.Second)
	}
}
