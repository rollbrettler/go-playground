package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/libgit2/git2go"
)

var clone string = "./test"
var repo string = "https://github.com/rollbrettler/go-playground.git"

func main() {

	cloneOptions := &git.CloneOptions{
		Bare: true,
	}

	_, err := ioutil.ReadDir(clone)
	if err != nil {
		log.Println(err)

		repository, err := git.Clone(repo, clone, cloneOptions)
		if err != nil {
			log.Println(err)
			return
		}
		config, err := repository.Config()
		if err != nil {
			log.Println(err)
			return
		}
		fmt.Printf("%v\n", config)
		return
	}
	log.Println("Folder already cloned")

	repository, err := git.OpenRepository(clone)
	config, err := repository.Config()
	if err != nil {
		log.Println(err)
		return
	}
  fetch, err := config.LookupString("remote.origin.fetch")
  mirror, err := config.LookupString("remote.origin.mirror")
	fmt.Printf("%v %v\n", fetch, mirror)
	config.SetString("remote.origin.fetch", "+refs/*:refs/*")
	config.SetBool("remote.origin.mirror", true)
}
