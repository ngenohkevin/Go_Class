package main

import (
	"fmt"
)

// type Pair struct {
// 	Path string
// 	Hash string
// }

// func (p Pair) String() string {
// 	return fmt.Sprintf("Hash of %s is %s", p.Path, p.Hash)
// }

// type PairWithLength struct {
// 	Pair
// 	Length int
// }

// func (p PairWithLength) String() string {
// 	return fmt.Sprintf("Hash of %s is %s, length %d", p.Path, p.Hash, p.Length)
// }

// func (p Pair) Filename() string {
// 	return filepath.Base(p.Path)
// }

// type Filenamer interface {
// 	Filename() string
// }

// func main() {
// 	p := Pair{"/usr", "Oxfdfe"}

// 	var fn Filenamer = PairWithLength{Pair{"/usr/lib", "Oxdead"}, 133}

// 	fmt.Println(p)
// 	fmt.Println(fn)
// 	fmt.Println(fn.Filename())
// }

type Shape interface {
	Area() float64
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func PrintArea(s Shape) {
	fmt.Printf("Area: %f\n", s.Area())
}

func main() {
	circle := Circle{Radius: 5}
	rectangle := Rectangle{Width: 4, Height: 3}

	PrintArea(circle)    // Prints the area of the circle
	PrintArea(rectangle) // Prints the area of the rectangle

	area := circle.Area() // Calling a method of a concrete type
	fmt.Println(area)
}
