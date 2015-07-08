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
	Url  string
	Path string
}

var clone string = "./test"
var repo string = "https://gitlab.com/rollbrettler/go-playground.git"

func main() {

	// repos := []Repos{
	// 	{Url: "https://gitlab.com/rollbrettler/go-playground.git", Path: "./test"},
	// 	{Url: "https://github.com/rollbrettler/compress-videos.git", Path: "./test2"},
	// }
	//
	// for _, repo := range repos {
	// 	fmt.Printf("%s -> %s\n", repo.Url, repo.Path)
	// }

	os.Exit(0)

	_, err := ioutil.ReadDir(clone)
	if err != nil {
		log.Println("Directory not pressent")

		err := cloneRepository()
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}
	}

	log.Println("Folder already cloned")

	repository, err := git.OpenRepository(clone)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	err = updateRepository(*repository)
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

	repository, err := git.Clone(repo.Url, repo.Path, cloneOptions)
	if err != nil {
		log.Println(err)
		return
	}

	config, err := repository.Config()
	if err != nil {
		log.Println(err)
		return
	}

	err = repo.changeToMirrorConfig(*config)
	if err != nil {
		log.Println(err)
	}

	err = repo.updateRepository(*repository)
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

func (repo Repo) updateRepository(repository git.Repository) (err error) {

	remote, err := repository.LookupRemote("origin")
	defer remote.Free()

	//refspecs := []string{"+refs/*:refs/*"}
	err = remote.Fetch([]string{}, nil, "")
	if err != nil {
		return err
	}

	return nil
}
