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
	text := string(file)
	fmt.Println(text)
}
