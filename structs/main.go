package main

import (
	"fmt"
)

type album1 struct {
	title string
}

type album2 struct {
	title string
}

func main() {

	var a1 = album1{
		"The White album",
	}
	var a2 = album2{

		"The Black album",
	}

	a1 = album1(a2)

	fmt.Println(a1, a2)
}
