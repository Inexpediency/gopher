package interfaces

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"os"
	"sort"
	"text/tabwriter"
	"time"
)

//import (
//	"bytes"
//	"flag"
//	"fmt"
//	"io"
//	"sort"
//	"time"
//)


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


//package sort
//type Interface interface {
//	Len() int
//	Less(i, j int) bool 11 i,j - индексы элементов в последовательности
//Swap(i, j int)
//}
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

func TestSort() {
	s := StringSlice{"name1", "name3", "name2"}
	fmt.Println("Input: ", s)
	sort.Sort(s)
	// Std lib: sort.SortStrings(s) :3
	fmt.Println("Result: ", s)
}
