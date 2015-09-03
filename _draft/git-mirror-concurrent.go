package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/libgit2/git2go"
)

// type Repos struct {
// 	repo []Repo
// }

type Repo struct {
	URL        string
	Path       string
	repository *git.Repository
}

var clone = "./test"
var repo = "https://gitlab.com/rollbrettler/go-playground.git"

func main() {

	// repos := []Repos{
	// 	{Url: "https://gitlab.com/rollbrettler/go-playground.git", Path: "./test"},
	// 	{Url: "https://github.com/rollbrettler/compress-videos.git", Path: "./test2"},
	// }
	//
	// for _, repo := range repos {
	// 	fmt.Printf("%s -> %s\n", repo.URL, repo.Path)
	// }

	var err error

	repo := Repo{URL: "https://gitlab.com/rollbrettler/go-playground.git", Path: "./test.git"}

	err = repo.openRepository()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	err = repo.updateRepository()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	log.Println("Sucessfully updated")
	os.Exit(0)
}

func (repo Repo) cloneRepository() (err error) {

	cloneOptions := &git.CloneOptions{
		Bare: true,
	}

	repo.repository, err = git.Clone(repo.URL, repo.Path, cloneOptions)
	if err != nil {
		log.Println(err)
		return
	}

	config, err := repo.repository.Config()
	if err != nil {
		log.Println(err)
		return
	}

	err = repo.changeToMirrorConfig(*config)
	if err != nil {
		log.Println(err)
	}

	err = repo.updateRepository()
	if err != nil {
		log.Println(err)
	}

	return nil
}

func (repo Repo) changeToMirrorConfig(config git.Config) (err error) {

	fetch, err := config.LookupString("remote.origin.fetch")
	if err != nil {
		return err
	}

	mirror, err := config.LookupBool("remote.origin.mirror")
	if err != nil {
		return err
	}

	log.Println(fmt.Sprintf("Refspec: %v\n", fetch))
	config.SetString("remote.origin.fetch", "+refs/*:refs/*")

	log.Println(fmt.Sprintf("Mirror: %v\n", mirror))
	config.SetBool("remote.origin.mirror", true)

	return nil
}

func (repo Repo) updateRepository() (err error) {
	remotes := repo.repository.Remotes
	remote, err := remotes.Lookup("origin")
	defer remote.Free()

	err = remote.Fetch([]string{}, nil, "")
	if err != nil {
		return err
	}

	return nil
}

func (repo Repo) openRepository() (err error) {
	_, err = ioutil.ReadDir(repo.Path)
	if err != nil {
		log.Println("Directory not present")

		err := repo.cloneRepository()
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}
	}

	log.Println("Folder already cloned")

	repo.repository, err = git.OpenRepository(repo.Path)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	return err
}
