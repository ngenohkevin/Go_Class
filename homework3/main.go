package main

import (
	"fmt"
	"log"
	"net/http"
)

type dollars float32

func (d dollars) String() string {
	return fmt.Sprintf("$%.2f", d)
}

type database map[string]dollars

func (db database) list(w http.ResponseWriter, req *http.Request) {

}

func main() {
	db := database{
		"shoes": 50,
		"socks": 5,
	}

	http.HandleFunc("/list", db.list)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
