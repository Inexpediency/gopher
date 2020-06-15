package webworkers

import (
	"fmt"
	"github.com/ythosa/gobih/malbedro"
	"github.com/ythosa/gobih/surface"
	"github.com/ythosa/gobih/types"
	"log"
	"net/http"
	"strconv"
	"strings"
	"sync"

	"github.com/ythosa/gobih/lissajous"
)

var (
	mu       sync.Mutex
	reqCount int
)

// StartServerAll ...
func StartServerAll() {
	http.HandleFunc("/", handlerCounter)
	http.HandleFunc("/count", requestsCounter)
	http.HandleFunc("/lissajous", lissajousHandler)
	http.HandleFunc("/surface", surfaceHandler)
	http.HandleFunc("/malbedro", malbedroHandler)
	http.HandleFunc("/nutonpic", nutonHandler)
	http.HandleFunc("/githubIssues", githubIssuesHandler)

	log.Fatal(http.ListenAndServe(serverURL, nil))
}

func getRequestData(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)

	fmt.Fprint(w, "Headers:\n")
	for k, v := range r.Header {
		fmt.Fprintf(w, "\t[%q] = %q\n", k, v)
	}

	fmt.Fprintf(w, "Host: %q\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr: %q\n", r.RemoteAddr)

	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}

	fmt.Fprint(w, "Forms:\n")
	for k, v := range r.Form {
		fmt.Fprintf(w, "\t[%q] = %q\n", k, v)
	}
}

func handlerCounter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	reqCount++
	mu.Unlock()

	getRequestData(w, r)
}

func requestsCounter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "Requests count: %d", reqCount)
	mu.Unlock()
}

func lissajousHandler(w http.ResponseWriter, r *http.Request) {
	// Request example: http://localhost:8080/lissajous?cycles=5&res=0.001&size=500&nframes=128&delay=5

	cycles, err := strconv.ParseInt(r.URL.Query().Get("cycles"), 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Invalid 'cycles' property input")
	}

	size, err := strconv.ParseInt(r.URL.Query().Get("size"), 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Invalid 'size' property input")
	}

	nframes, err := strconv.ParseInt(r.URL.Query().Get("nframes"), 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Invalid 'nframes' property input")
	}

	delay, err := strconv.ParseInt(r.URL.Query().Get("delay"), 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Invalid 'delay' property input")
	}

	res, err := strconv.ParseFloat(r.URL.Query().Get("res"), 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Invalid 'res' property input")
	}

	lissajous.Draw(w, int(cycles), res, int(size), int(nframes), int(delay))
}

func surfaceHandler(w http.ResponseWriter, r *http.Request) {
	// Request example: http://localhost:8080/surface?width=300&height=500&cells=100&xyrange=30

	w.Header().Set("Content-Type", "image/svg+xml")

	width, err := strconv.ParseInt(r.URL.Query().Get("width"), 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Invalid 'width' property input")
	}

	height, err := strconv.ParseInt(r.URL.Query().Get("height"), 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Invalid 'height' property input")
	}

	cells, err := strconv.ParseInt(r.URL.Query().Get("cells"), 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Invalid 'cells' property input")
	}

	xyrange, err := strconv.ParseFloat(r.URL.Query().Get("xyrange"), 10)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Invalid 'xyrange' property input")
	}

	s := surface.Surf{
		Width: int(width),
		Height: int(height),
		Cells: int(cells),
		XYRange: xyrange,
	}

	s.Draw(w)
}

func malbedroHandler(w http.ResponseWriter, r *http.Request) {
	malbedro.DrawMalbedro(w)
}

func nutonHandler(w http.ResponseWriter, r *http.Request) {
	malbedro.DrawNuton(w)
}

func githubIssuesHandler(w http.ResponseWriter, r *http.Request) {
	// Request example: http://localhost:8080/githubIssues?repo:ythosa/where-is
	args := strings.Split(r.URL.String()[14:], "&")
	types.ServerSearch(w, args)
}
