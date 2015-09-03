package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/libgit2/git2go"
)

var clone string = "./test.git"
var repo string = "https://github.com/libgit2/git2go.git"

func main() {

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

func cloneRepository() (err error) {

	cloneOptions := &git.CloneOptions{
		Bare: true,
	}

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
	if err != nil {
		log.Println(err)
	}

	err = updateRepository(*repository)
	if err != nil {
		log.Println(err)
	}

	fmt.Printf("%v\n", &config)

	return nil
}

func changeToMirrorConfig(config git.Config) (err error) {

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

func updateRepository(repository git.Repository) (err error) {

	remote, err := repository.Remotes.Lookup("origin")
	defer remote.Free()

	refspecs := []string{"+refs/*:refs/*"}
	err = remote.Fetch(refspecs, nil, "")
	if err != nil {
		return err
	}

	return nil
}
