package interfaces

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

/*
package http
type Handler interface {
	ServeHTTP(w ResponseWriter, r *Request)
}
func ListenAndServe(address string, h Handler) error
*/

type dollars float32

func (d dollars) String() string {
	return fmt.Sprintf("$%.2f", d)
}

type database map[string]dollars

/* Shop 1.0 */

// OpenShop starts shop server
func OpenShop() {
	db := database{"shoes": 50, "socks": 30}
	log.Fatal(http.ListenAndServe("localhost:8081", db))
}

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
	case "/setprice":
		item := r.URL.Query().Get("item")
		price, err := strconv.ParseFloat(r.URL.Query().Get("price"), 10)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, "Invalid request")
			return
		}

		lastPrice, ok := db[item]
		db[item] = dollars(price)
		if ok {
			fmt.Fprintf(w, "Price of %s was changed from %s to %s\n", item, lastPrice, db[item])
		} else {
			fmt.Fprintf(w, "Added new item: %s with cost = %s", item, db[item])
		}
	default:
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "There is not this page: %s", r.URL)
	}
}

/* Shop 2.0 */

// OpenBetterShop starts better shop server
func OpenBetterShop() {
	db := database{"shoes": 50, "socks": 30}
	mux := http.NewServeMux()
	mux.HandleFunc("/list", db.list)
	mux.HandleFunc("/price", db.price)
	mux.HandleFunc("/setprice", db.setPrice)
	log.Fatal(http.ListenAndServe("localhost:8081", mux))
}

func (db database) list(w http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

func (db database) price(w http.ResponseWriter, r *http.Request) {
	item := r.URL.Query().Get("item")
	price, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "There is no this item: %q\n", item)
		return
	}
	fmt.Fprintf(w, "Cost of %s = %s\n", item, price)
}

func (db database) setPrice(w http.ResponseWriter, r *http.Request) {
	item := r.URL.Query().Get("item")
	price, err := strconv.ParseFloat(r.URL.Query().Get("price"), 10)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Invalid request")
		return
	}

	lastPrice, ok := db[item]
	db[item] = dollars(price)
	if ok {
		fmt.Fprintf(w, "Price of %s was changed from %s to %s\n", item, lastPrice, db[item])
	} else {
		fmt.Fprintf(w, "Added new item: %s with cost = %s", item, db[item])
	}
}

/* Shop 3.0 */

// OpenTheBestShop starts the best shop server realisation
func OpenTheBestShop() {
	db := database{"shoes": 50, "socks": 30}
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/price", db.price)
	http.HandleFunc("/setprice", db.setPrice)

	// A global ServeMux instance named DefaultServeMux and the
	// http.Handle and http package level functions http.HandleFunc
	log.Fatal(http.ListenAndServe("localhost:8081", nil))
}
