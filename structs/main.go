package main

import (
	"fmt"
)

func main() {

	var album = struct {
		title  string
		artist string
		year   int
		copies int
	}{
		"The White album",
		"The Beatles",
		1968,
		1000000000000,
	}

	var pAlbum = &album

	fmt.Println(album, pAlbum)
}
