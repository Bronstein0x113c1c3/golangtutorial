package main

import (
	"chap7/components"
	"log"
	"net/http"
)

func main() {
	db := components.Database{"shoes": 50, "socks": 5}
	http.HandleFunc("/list", db.List)
	http.HandleFunc("/price", db.Price)
	http.HandleFunc("/", db.Entry)
	http.HandleFunc("/create", (&db).Create)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))

}
