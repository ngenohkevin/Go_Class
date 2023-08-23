package main

import "fmt"

type Organ struct {
	Name   string
	Weight int
}

type Organs []Organ

func (s Organs) Len() int {
	return len(s)
}
func (s Organs) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

type ByName struct {
	Organs
}
type ByWeight struct {
	Organs
}

func main() {
	s := []Organ{{"brain", 1340}, {"liver", 1494}, {"spleen", 162}, {"pancreas", 131}, {"heart", 290}}

	fmt.Println(s)
}
