package main

import "fmt"

func main() {
	var s []int

	fmt.Printf("%d, %d, %T, %5t, %#[3]v\n", len(s), cap(s), s, s == nil)
}
