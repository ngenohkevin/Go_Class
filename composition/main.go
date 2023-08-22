package main

import "fmt"

type Pair struct {
	Path string
	Hash string
}

func (p Pair) String() string {
	return fmt.Sprintf("Hash of %s is %s", p.Path, p.Hash)
}

func main() {
	p := Pair{"/usr", "Oxfdfe"}

	fmt.Println(p)
}
