package main

import "time"

type T struct {
	i byte
	b bool
}

func send(i int, ch chan<- *T) {
	t := &T{i: byte(i)}
	ch <- t

	t.b = true //Unsafe at any speed
}

func main() {
	vs := make([]T, 5)
	ch := make(chan *T)

	for i := range vs {
		go send(i, ch)
	}

	time.Sleep(1 * time.Second) //all goroutines started
}
