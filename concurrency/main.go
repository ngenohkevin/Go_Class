package main

import (
	"log"
	"net/http"
	"time"
)

type result struct {
	url     string
	err     error
	latency time.Duration
}

func get(url string, ch chan<- result) {
	start := time.Now()

	if resp, err := http.Get(url); err != nil {
		ch <- result{url, err, 0}
	} else {
		t := time.Since(start).Round(time.Millisecond)
		ch <- result{url, nil, t}
		resp.Body.Close()
	}
}

func main() {

	stopper := time.After(4 * time.Second)

	results := make(chan result)
	list := []string{
		"https://amazon.com",
		"https://www.arnoderrymovers.co.ke",
		"https://google.com",
		"https://github.com",
		"https://nytimes.com",
		"http://wsj.com",
	}

	for _, url := range list {
		go get(url, results)
	}

	for range list {

		select {
		case r := <-results:
			if r.err != nil {
				log.Printf("%-20s %s\n", r.url, r.err)
			} else {
				log.Printf("%-20s %s\n", r.url, r.latency)
			}
		case t := <-stopper:
			log.Fatalf("timeout %s", t)
		}
	}
}

// func generator(limit int, ch chan<- int) {
// 	for i := 2; i < limit; i++ {
// 		ch <- i
// 	}
// 	close(ch)
// }

// func filter(src <-chan int, dst chan<- int, prime int) {
// 	for i := range src {
// 		if i%prime != 0 {
// 			dst <- i
// 		}
// 	}

// 	close(dst)
// }

// func sieve(limit int) {
// 	ch := make(chan int)

// 	go generator(limit, ch)

// 	for {
// 		prime, ok := <-ch

// 		if !ok {
// 			break
// 		}
// 		ch1 := make(chan int)

// 		go filter(ch, ch1, prime)

// 		ch = ch1

// 		fmt.Print(prime, " ")
// 	}
// 	fmt.Println()
// }
// func main() {
// 	sieve(100)
// }
