package main

import (
	"strconv"
	"strings"
)

type IntSlice []int

func (is IntSlice) String() string {

	var strs []string

	for _, v := range is {
		strs = append(strs, strconv.Itoa(v))
	}

	return "[" + strings.Join(strs, ";") + "]"
}
