package main

import "fmt"

type Pair struct {
	Path string
	Hash string
}

func (p Pair) String() string {
	return fmt.Sprintf("Hash of %s is %s", p.Path, p.Hash)
}

type PairWithLength struct {
	Pair
	Length int
}

func main() {
	p := Pair{"/usr", "Oxfdfe"}

	pl := PairWithLength{Pair{"/usr/lib", "Oxdead"}, 133}

	fmt.Println(p)
	fmt.Println(pl)
}
