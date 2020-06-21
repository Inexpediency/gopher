package interfaces

import (
	"fmt"
	"os"
	"sort"
	"text/tabwriter"
	"time"
)

// Track type
type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

var tracks = []*Track{
	{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
	{"Go", "Moby", "Moby", 1992, length("3m37s")},
	{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
	{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
}

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

func printTracks(tracks []*Track) {
	const format = "%v\t%v\t%v\t%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Title", "Artist", "Album", "Year", "Length")
	fmt.Fprintf(tw, format, "---- ", "------", "-----", "----", "----- ")
	for _, t := range tracks {
		fmt.Fprintf(tw, format, t.Title, t.Artist,
			t.Album, t.Year, t.Length)
	}
	tw.Flush() // Counting columns size and output table
}

type customSort struct {
	t    []*Track
	less func(x, y *Track) bool
}

func (x customSort) Len() int           { return len(x.t) }
func (x customSort) Less(i, j int) bool { return x.less(x.t[i], x.t[j]) }
func (x customSort) Swap(i, j int)      { x.t[i], x.t[j] = x.t[j], x.t[i] }

// SortTasks function
func SortTasks(t string, rev bool) error {
	if t == "all" {
		sort.Sort(customSort{tracks, func(x, y *Track) bool {
			if x.Title != y.Title {
				return x.Title < y.Title
			}
			if x.Year != y.Year {
				return x.Year < y.Year
			}
			if x.Length != y.Length {
				return x.Length < y.Length
			}
			return false
		}})
	} else if t == "artist" {
		if rev {
			sort.Sort(customSort{tracks, func(x, y *Track) bool {
				return x.Artist < y.Artist
			}})
		} else {
			sort.Sort(sort.Reverse(customSort{tracks, func(x, y *Track) bool {
				return x.Artist < y.Artist
			}}))
		}
	} else if t == "year" {
		if rev {
			sort.Sort(customSort{tracks, func(x, y *Track) bool {
				return x.Year < y.Year
			}})
		} else {
			sort.Sort(sort.Reverse(customSort{tracks, func(x, y *Track) bool {
				return x.Year < y.Year
			}}))
		}
	} else if t == "length" {
		if rev {
			sort.Sort(customSort{tracks, func(x, y *Track) bool {
				return x.Length < y.Length
			}})
		} else {
			sort.Sort(sort.Reverse(customSort{tracks, func(x, y *Track) bool {
				return x.Length < y.Length
			}}))
		}
	} else {
		return fmt.Errorf("invalid tasks sort type")
	}

	printTracks(tracks)
	return nil
}
