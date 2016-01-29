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
	for index, element := range folder {
		var dir string
		if(element.IsDir()) {
			dir = "📁 "
		}
		fmt.Printf("%d %v %v\n", index, dir, element.Name())
	}
}
