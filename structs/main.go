package main

import "fmt"

type Response struct {
	Page  int      `json:"page"`
	Words []string `json:"words,omitempty"`
}

func main() {
	r := &Response{Page: 1, Words: []string{"up", "in", "out"}}

	fmt.Printf("%#v\n", r)
}
