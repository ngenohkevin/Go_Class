package main

import "time"

type result struct {
	url     string
	err     error
	latency time.Duration
}
