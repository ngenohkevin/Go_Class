package main

import (
	"fmt"
	"time"
)

type Employee struct {
	Name   string
	Number int
	Boss   *Employee
	Hired  time.Time
}

func main() {

	c := map[string]*Employee{}

	c["Dilan"] = &Employee{"Dilan", 2, nil, time.Now()}

	c["Kevin"] = &Employee{
		Name:   "Kevin",
		Number: 1,
		Boss:   c["Lamine"],
		Hired:  time.Now(),
	}

	fmt.Printf("%T %+[1]v\n", c["Dilan"])
	fmt.Printf("%T %+[1]v\n", c["Kevin"])
}
