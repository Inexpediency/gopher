package interfaces

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"time"
)


type ByteCounter int
func (c *ByteCounter) Write (p []byte) (int, error) {
	*c += ByteCounter(len(p))
	return len(p), nil
}


// *bytes.Buffer должен соответствовать io.Writer
var _ io.Writer = (*bytes.Buffer)(nil)


type Artifact interface {
	Title() string
	Creators() []string
	Created() time.Time
}
type Text interface {
	Pages() int
	Words() int
	PageSize() int
}
type Audio interface {
	Stream() (io.ReadCloser, error)
	RunningTime() time.Duration
	Format() string // Например, "MP3", "WAV"
}
type Video interface {
	Stream() (io.ReadCloser, error)
	RunningTime() time.Duration
	Format() string // Например, "MP4", "WMV"
	Resolution() (x, у int)
}
type Streamer interface {
	Stream() (io.ReadCloser, error)
	RunningTime() time.Duration
	Format() string
}


func Waiter() {
	// build example main --period 10s
	var period = flag.Duration("period", 1*time.Second, "sleep period")
	flag.Parse()
	fmt.Printf("Ожидание %v...", *period)
	time.Sleep(*period)
	fmt.Println()
}
