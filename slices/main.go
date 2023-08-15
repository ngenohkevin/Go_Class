package main

import "fmt"

func main() {
	a := [3]int{1, 2, 3}
	b := a[:1]

	fmt.Println("a == ", a)
	fmt.Println("b == ", b)

	c := b[:2]
	fmt.Println("c == ", c)

}
