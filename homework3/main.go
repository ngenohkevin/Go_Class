package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type dollars float32

func (d dollars) String() string {
	return fmt.Sprintf("$%.2f", d)
}

type database map[string]dollars

func (db database) list(w http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s : %s\n", item, price)
	}
}

func (db database) add(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price := req.URL.Query().Get("price")

	if _, ok := db[item]; ok {
		msg := fmt.Sprintf("duplicate item: %q", item)
		http.Error(w, msg, http.StatusBadRequest)
		return
	}

	p, err := strconv.ParseFloat(price, 32)
	if err != nil {
		msg := fmt.Sprintf("Invalid price: %q", price)
		http.Error(w, msg, http.StatusBadRequest)
		return
	}

	db[item] = dollars(p)

	fmt.Fprintf(w, "added %s with price %s\n", item, db[item])
}
func (db database) Update(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price := req.URL.Query().Get("price")

	if _, ok := db[item]; !ok {
		msg := fmt.Sprintf("no such item: %q", item)
		http.Error(w, msg, http.StatusNotFound)
		return
	}

	p, err := strconv.ParseFloat(price, 32)
	if err != nil {
		msg := fmt.Sprintf("Invalid price: %q", price)
		http.Error(w, msg, http.StatusBadRequest)
		return
	}

	db[item] = dollars(p)

	fmt.Fprintf(w, "added %s with price %s\n", item, db[item])
}

func main() {
	db := database{
		"shoes": 50,
		"socks": 5,
	}

	http.HandleFunc("/list", db.list)
	http.HandleFunc("/create", db.add)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
