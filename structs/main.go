package main

import "time"

type Employee struct {
	Name   string
	Number int
	Boss   *Employee
	Hired  time.Time
}
