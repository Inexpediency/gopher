package interfaces

import (
	"fmt"
	"log"
	"net/http"
)

/*
package http
type Handler interface {
	ServeHTTP(w ResponseWriter, r *Request)
}
func ListenAndServe(address string, h Handler) error
*/

func OpenShop() {
	db := database{"shoes": 50, "socks": 30}
	log.Fatal(http.ListenAndServe("localhost:8081", db))
}

type dollars float32

func (d dollars) String() string {
	return fmt.Sprintf("$%.2f", d)
}

type database map[string]dollars

func (db database) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/list":
		for item, price := range db {
			fmt.Fprintf(w, "%s: %s\n", item, price)
		}
	case "/price":
		item := r.URL.Query().Get("item")
		price, ok := db[item]
		if !ok {
			w.WriteHeader(http.StatusNotFound) // 404
			fmt.Fprintf(w, "There is no this item: %q\n", item)
			return
		}
		fmt.Fprintf(w, "Cost of %s = %s\n", item, price)
	default:
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "There is not this page: %s", r.URL)
	}
}
