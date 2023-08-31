package main

import (
	"net/http"
	"time"
)

type result struct {
	url     string
	err     error
	latency time.Duration
}

func get(url, ch chan<- result) {
	start := time.Now()

	if resp, err := http.Get(url); err != nil {

	} else {
		resp.Body.Close()
	}
}
