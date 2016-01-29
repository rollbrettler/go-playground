package main

import (
	"fmt"
	"strconv"
)

func main() {
	var i int
	var s string

	i = 1
	s = strconv.Itoa(i)

	fmt.Printf("%T: %v\n", s, s)
}
