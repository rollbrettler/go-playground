package main

import (
	"fmt"
	"log"

	"github.com/libgit2/git2go"
)

func showCommitId(commit *git.Commit) bool {
	fmt.Printf("%b\n", commit.Message())
	return true
}

func main() {
	repo, err := git.OpenRepository(".")
	if err != nil {
		log.Println(err)
	}

	h, _ := repo.Head()
	n, _ := h.Branch().Name()
	fmt.Printf("Branch name: %v\n", n)

	i, _ := repo.Index()
	e := i.EntryCount()
	fmt.Printf("Files in index: %v\n", e)

  id, _ := headCommitId(repo)
	fmt.Printf("Head id: %v\n", id)

	// walk, _ := repo.Walk()
	// walk.Iterate(showCommitId)
}

// https://github.com/pjvds/pressurizer/blob/5c48a2c630c66526e0a9a00a7c879e1c76e82be6/GitWatcher.go#L80
func headCommitId(repo *git.Repository) (*git.Oid, error) {
	headRef, err := repo.LookupReference("HEAD")
	defer headRef.Free()
	if err != nil {
		return nil, err
	}

	ref, err := headRef.Resolve()
	defer ref.Free()
	if err != nil {
		return nil, err
	}

	return ref.Target(), nil
}
