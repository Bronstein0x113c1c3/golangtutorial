package components

import (
	"fmt"
	"net/http"
	"strconv"
	"sync"
)

var mu sync.Mutex

func (db Database) List(w http.ResponseWriter, r *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s \n", item, price)
	}
}

func (db Database) Price(w http.ResponseWriter, r *http.Request) {
	item := r.URL.Query().Get("item")
	price, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "no such item: %q", item)
		return
	}
	fmt.Fprintf(w, "%s: %s \n", item, price)
}
func (db Database) Entry(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to our store")
	db.List(w, r)
}
func (db *Database) Create(w http.ResponseWriter, r *http.Request) {
	item := r.URL.Query().Get("item")
	if item == "" {
		fmt.Fprint(w, "Item not provided")
		return
	}
	price_for_newitem := r.URL.Query().Get("price")
	if price_for_newitem == "" {
		fmt.Fprintln(w, "Price not provided")
		return
	}
	_, ok := (*db)[item]
	if !ok {
		price, err := strconv.ParseFloat(price_for_newitem, 64)
		if err != nil {
			http.Error(w, "Error in parsing number", http.StatusUnprocessableEntity)
			return
		}
		mu.Lock()
		(*db)[item] = Dollars(price)
		mu.Unlock()
		fmt.Fprintln(w, "Created successfully!")
		return
	}
	fmt.Fprint(w, "This item is found, so cannot created!")
	return
}
