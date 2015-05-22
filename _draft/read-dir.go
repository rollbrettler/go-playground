package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	folder, err := ioutil.ReadDir(".")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(folder)
}
