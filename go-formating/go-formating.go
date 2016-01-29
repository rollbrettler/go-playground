package main

import "fmt"

type polar struct {
	radius float64
	θ      float64
}

func main() {
	customType := polar{2.2, 2.3}
	fmt.Printf("Default output: %v\n", customType)
	fmt.Printf("Alternative output: %#v\n", customType)
	fmt.Printf("%20s\n", "Test string")
	fmt.Printf("%-20s\n", "Test string")
}
