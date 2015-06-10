package main

import (
	_ "fmt"
	"io/ioutil"
	"log"

	"github.com/libgit2/git2go"
)

var clone string = "./test"
var repo string = "https://github.com/rollbrettler/go-playground.git"

func main() {

	var cloneOptions git.CloneOptions
	cloneOptions.Bare = true

	_, err := ioutil.ReadDir(clone)
	if err != nil {
		log.Println(err)

		_, err := git.Clone(repo, clone, &cloneOptions)
		if err != nil {
			log.Println(err)
			return
		}
		return
	}
	log.Println("Folder already cloned")
}
