package interfaces

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"sort"
	"time"
)

// ByteCounter type
type ByteCounter int

func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p))
	return len(p), nil
}

// *bytes.Buffer must match io.Writer
var _ io.Writer = (*bytes.Buffer)(nil)

// Artifact type
type Artifact interface {
	Title() string
	Creators() []string
	Created() time.Time
}

// Text type
type Text interface {
	Pages() int
	Words() int
	PageSize() int
}

// Audio type
type Audio interface {
	Stream() (io.ReadCloser, error)
	RunningTime() time.Duration
	Format() string // For example, "MP3", "WAV"
}

// Video type
type Video interface {
	Stream() (io.ReadCloser, error)
	RunningTime() time.Duration
	Format() string // For example, "MP4", "WMV"
	Resolution() (x, у int)
}

// Streamer type
type Streamer interface {
	Stream() (io.ReadCloser, error)
	RunningTime() time.Duration
	Format() string
}

// Waiter CLI application
func Waiter() {
	// build example:  main --period 10s
	var period = flag.Duration("period", 1*time.Second, "sleep period")
	flag.Parse()
	fmt.Printf("Ожидание %v...", *period)
	time.Sleep(*period)
	fmt.Println()
}

//package sort
//type Interface interface {
//	Len() int
//	Less(i, j int) bool 11 i,j - индексы элементов в последовательности
//Swap(i, j int)
//}

// StringSlice type
type StringSlice []string

func (p StringSlice) Len() int {
	return len(p)
}

func (p StringSlice) Less(i, j int) bool {
	return p[i] < p[j]
}

func (p StringSlice) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

// TestSort represents sort example
func TestSort() {
	s := StringSlice{"name1", "name3", "name2"}
	fmt.Println("Input: ", s)
	sort.Sort(s)
	// Std lib: sort.SortStrings(s) :3
	fmt.Println("Result: ", s)
}

/* Using type declaration example */
func highloadWriteHeader(w io.Writer, contentType string) error {
	/* The most highload server part xd */
	if _, err := writeString(w, "Content-Type: "); err != nil {
		return err
	}
	if _, err := writeString(w, contentType); err != nil {
		return err
	}
	// ...
	return nil
}

// WriteString writes `s` to `w`
// If `w` has the `WriteString` method, it is called instead of `w.Write`
func writeString(w io.Writer, s string) (n int, err error) {
	type stringWriter interface {
		WriteString(string) (n int, err error)
	}
	if sw, ok := w.(stringWriter); ok {
		return sw.WriteString(s) // Avoid copy string
	}
	return w.Write([]byte(s)) // Using temporary copy
}

// Type definition with switch example
func sqlQuote(x interface{}) string {
	switch x := x.(type) {
	case nil:
		return "NULL"
	case int, uint:
		return fmt.Sprintf("%d", x) // X has type interface{}.
	case bool:
		if x {
			return "TRUE"
		}
		return "FALSE"
	// string and other types ...
	default:
		panic(fmt.Sprintf("unexpected type %T: %v", x, x))
	}
}
