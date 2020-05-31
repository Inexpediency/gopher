package webworkers

import (
	"fmt"
	"log"
	"net/http"
)

const serverURL = "localhost:8080"

// StartServer ...
func StartServer() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(serverURL, nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}
