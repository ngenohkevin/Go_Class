package main

import (
	"fmt"
)

type album1 struct {
	title string
}

func main() {

	var album1 = struct {
		title string
	}{
		"The White album",
	}
	var album2 = struct {
		title string
	}{
		"The Black album",
	}

	// album1 = album2

	fmt.Println(album1, album2)
}
