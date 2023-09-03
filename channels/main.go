package main

import "fmt"

func main() {
	ch := make(chan int)

	ch <- 1
	b, ok := <-ch
	fmt.Println(b, ok)

	close(ch)

	c, ok := <-ch //he
	fmt.Println(c, ok)
}
