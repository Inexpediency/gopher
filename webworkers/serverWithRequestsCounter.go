package webworkers

import (
	"fmt"
	"github.com/ythosa/gobih/malbedro"
	"github.com/ythosa/gobih/surface"
	"log"
	"math"
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

// StartServerRequestsCounter ...
func StartServerRequestsCounter() {
	http.HandleFunc("/", handlerCounter)
	http.HandleFunc("/count", requestsCounter)
	http.HandleFunc("/lissajous", lissajousHandler)
	http.HandleFunc("/surface", surfaceHandler)
	http.HandleFunc("/malbedro", malbedroHandler)
	http.HandleFunc("/nutonpic", nutonHandler)

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
	// http://localhost:8080/lissajous?cycles=5&res=0.001&size=500&nframes=128&delay=5

	var cycles, size, nframes, delay int
	var res float64

	url := strings.Split(r.URL.String()[11:], "&")
	for _, v := range url {
		prop := strings.Split(v, "=")[0]
		value := strings.Split(v, "=")[1]

		if prop == "cycles" {
			cycles, _ = strconv.Atoi(value)
		} else if prop == "res" {
			res, _ = strconv.ParseFloat(value, 64)
		} else if prop == "size" {
			size, _ = strconv.Atoi(value)
		} else if prop == "nframes" {
			nframes, _ = strconv.Atoi(value)
		} else if prop == "delay" {
			delay, _ = strconv.Atoi(value)
		}
	}

	fmt.Printf("res: %f, cycles: %d, size: %d, nframes: %d, delay: %d\n", res, cycles, size, nframes, delay)

	lissajous.Draw(w, cycles, res, size, nframes, delay)
}

func surfaceHandler(w http.ResponseWriter, r *http.Request) {
	// http://localhost:8080/surface?width=300&height=500&cells=100&xyrange=30

	w.Header().Set("ContentType", "image/svg+xml")

	var (
		width, height, cells int
		xyrange, xyscale, zscale, angle float64
		)

	url := strings.Split(r.URL.String()[9:], "&")
	for _, v := range url {
		prop := strings.Split(v, "=")[0]
		value := strings.Split(v, "=")[1]

		if prop == "width" {
			width, _ = strconv.Atoi(value)
		} else if prop == "height" {
			height, _ = strconv.Atoi(value)
		} else if prop == "cells" {
			cells, _ = strconv.Atoi(value)
		} else if prop == "xyrange" {
			xyrange, _ = strconv.ParseFloat(value, 64)
		}
	}

	xyscale = float64(width)/2/xyrange
	zscale = float64(height) * 0.4
	angle = math.Pi / 6

	s := surface.Surf{Width: width, Height: height, Cells: cells, XYrange: xyrange, XYscale: xyscale, Zscale: zscale, Angle: angle}

	surface.Draw(w, &s)
}

func malbedroHandler(w http.ResponseWriter, r *http.Request) {
	malbedro.DrawMalbedro(w)
}

func nutonHandler(w http.ResponseWriter, r *http.Request) {
	malbedro.DrawNuton(w)
}