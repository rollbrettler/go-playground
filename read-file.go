package main

import (
	"fmt"
	_ "io/ioutil"
	"os"
)

func main() {
	//f, err := ioutil.ReadFile("README.md")
	f, err := os.Open("README.md")
	if err != nil {
		return
	}
	fmt.Println(f.Name())
}
