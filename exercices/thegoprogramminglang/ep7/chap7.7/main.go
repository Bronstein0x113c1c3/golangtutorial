package main

import (
	"fmt"
	"log"
	"net/http"
)

// the price for each type of goods
type dollars float64

func (d dollars) String() string {
	return fmt.Sprintf("$%.2f", d)
}

//the dollars type implements fmt.Stringer interface, so when we use some method including:

// the list of the goods and also, its priceS
type database map[string]dollars

// but just the ServeHTTP is too complex.....

// func (db database) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	//check the url path
// 	switch r.URL.Path {
// 	case "/list":
// 		for good, price := range db {
// 			fmt.Fprintf(w, "%s:%s \n", good, price)
// 		}
// 	case "/price":
// 		item := r.URL.Query().Get("item")
// 		// find the name of item.
// 		price, found := db[item]
// 		if !found {
// 			w.WriteHeader(http.StatusNotFound)
// 			fmt.Fprintf(w, "no such item: %q \n", item)
// 			return
// 		}
// 		fmt.Fprintf(w, "%v: %v \n", item, price)
// 	default:
// 		msg := fmt.Sprintf("no such page: %v \n", r.URL)
// 		http.Error(w, msg, http.StatusNotFound)
// 	}

// }

//divide into 2 function: list/price

func (db database) list(w http.ResponseWriter, r *http.Request) {
	for good, price := range db {
		fmt.Fprintf(w, "%s:%s \n", good, price)
	}
}
func (db database) price(w http.ResponseWriter, r *http.Request) {

	item := r.URL.Query().Get("item")
	price, found := db[item]
	if !found {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such item: %q \n", item)
		return
	}
	fmt.Fprintf(w, "%v: %v \n", item, price)
}

func main() {
	db := database{"shoes": 50, "socks": 5}
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/price", db.price)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
