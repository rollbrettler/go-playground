package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	file, err := ioutil.ReadFile("README.md")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// Convert utf-8 binary code into utf-8 strings
	text := string(file)
	fmt.Println(text)
}
