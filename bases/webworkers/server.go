package webworkers

import (
	"fmt"
	"log"
	"net/http"
)

// StartServer ...
func StartServer() {
	const url = "localhost:8080"

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(url, nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}
