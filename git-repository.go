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
  fmt.Println(h.Branch().Name())
}
