package main

import (
	"fmt"
	"log"

  "github.com/libgit2/git2go"
)

func main() {
  repo, err := git.OpenRepository(".")
  if err != nil {
    log.Println(err)
  }
  h, _ := repo.Head()
  n, _ := h.Branch().Name()
  fmt.Printf("branch name: %v\n", n)

  i, _ := repo.Index()
  e := i.EntryCount()
  fmt.Printf("index: %v\n", e)
}
