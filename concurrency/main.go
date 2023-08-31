package main

import (
	"fmt"
	"log"
	"net/http"
)

// type result struct {
// 	url     string
// 	err     error
// 	latency time.Duration
// }

// func get(url string, ch chan<- result) {
// 	start := time.Now()

// 	if resp, err := http.Get(url); err != nil {
// 		ch <- result{url, err, 0}
// 	} else {
// 		t := time.Since(start).Round(time.Millisecond)
// 		ch <- result{url, nil, t}
// 		resp.Body.Close()
// 	}
// }

// func main() {
// 	results := make(chan result)
// 	list := []string{
// 		"https://amazon.com",
// 		"https://www.arnoderrymovers.co.ke",
// 		"https://google.com",
// 		"https://github.com",
// 		"https://nytimes.com",
// 		"http://wsj.com",
// 	}

// 	for _, url := range list {
// 		go get(url, results)
// 	}

// 	for range list {
// 		r := <-results

// 		if r.err != nil {
// 			log.Printf("%-20s %s\n", r.url, r.err)
// 		} else {
// 			log.Printf("%-20s %s\n", r.url, r.latency)
// 		}
// 	}
// }

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1> You got %d</h1>", nextID)
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
