package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/libgit2/git2go"
)

var clone string = "./test"
var repo string = "https://gitlab.com/rollbrettler/go-playground.git"

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
		err = changeToMirrorConfig(*config)
		if (err != nil) {
			log.Println(err)
		}
		err = updateRepository(*repository)
		if (err != nil) {
			log.Println(err)
		}
		fmt.Printf("%v\n", config)
		return
	}
	log.Println("Folder already cloned")

	repository, err := git.OpenRepository(clone)
	if err != nil {
		log.Println(err)
		return
	}
	err = updateRepository(*repository)
	if (err != nil) {
		log.Println(err)
	}
}

func changeToMirrorConfig(config git.Config) (err error) {
	fetch, err := config.LookupString("remote.origin.fetch")
	if (err != nil) {
		log.Println(err)
	}
	mirror, err := config.LookupString("remote.origin.mirror")
	if (err != nil) {
		log.Println(err)
	}
	fmt.Printf("%v %v\n", fetch, mirror)
	config.SetString("remote.origin.fetch", "+refs/*:refs/*")
	config.SetBool("remote.origin.mirror", true)
	return nil
}

func updateRepository(repository git.Repository) (err error) {
	remote, err := repository.LookupRemote("origin")
	defer remote.Free()
	err = remote.FetchRefspecs([]string{}, nil, "")
	if(err != nil){
		return err
	}
	return nil
}
