package main

import (
	"fmt"

	"github.com/libgit2/git2go"
)

func printCommitMessages(commit *git.Commit) bool {
	fmt.Printf("Commit message: %v", commit.Message())
	return true
}

func main() {
	repo, err := git.OpenRepository(".")
	if err != nil {
    fmt.Printf("%v\n",err)
	}

	walk, err := repo.Walk()
	if err != nil {
    fmt.Printf("%v\n",err)
	}
  // Defer gets executet after the function is finished
  defer walk.Free()

  // Push the range the walker should iterate
  walk.PushRange("HEAD~2..HEAD")
	walk.Iterate(printCommitMessages)
}
