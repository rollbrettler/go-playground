package main

import (
	_ "fmt"
	"log"

	"github.com/libgit2/git2go"
)

func main() {

  var cloneOptions git.CloneOptions
  cloneOptions.Bare = true

  if _, err := git.Clone("https://github.com/rollbrettler/go-playground.git", "./test", &cloneOptions); err != nil {
		log.Println(err)
	}
}
