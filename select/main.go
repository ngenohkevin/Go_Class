package main

import (
	"log"
	"time"
)

func main() {
	chans := []chan int{
		make(chan int),
		make(chan int),
	}

	for i := range chans {
		go func(i int, ch chan<- int) {
			for {
				time.Sleep(time.Duration(i) * time.Second)
				ch <- i
			}
		}(i+1, chans[i])
	}

	for i := 0; 1 < 12; i++ {
		select {
		case m0 := <-chans[0]:
			log.Println("recieved", m0)
		case m1 := <-chans[1]:
			log.Println("recieve", m1)
		}
	}
}
