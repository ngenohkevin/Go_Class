package main

import (
	"fmt"
	"path/filepath"
)

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

func (p PairWithLength) String() string {
	return fmt.Sprintf("Hash of %s is %s, length %d", p.Path, p.Hash, p.Length)
}

func Filename(p Pair) string {
	return filepath.Base(p.Path)
}
func main() {
	p := Pair{"/usr", "Oxfdfe"}

	pl := PairWithLength{Pair{"/usr/lib", "Oxdead"}, 133}

	fmt.Println(p)
	fmt.Println(pl)
	fmt.Println(Filename(p))
}
