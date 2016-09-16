package main

import (
	"fmt"
	"os"
	"runtime"

	"golang.org/x/oauth2"

	"github.com/google/go-github/github"
)

var token string
var workers = runtime.NumCPU()

func main() {
	if token = os.Getenv("TOKEN"); token == "" {
		fmt.Printf("Please specify a TOKEN\n")
		os.Exit(1)
	}

	runtime.GOMAXPROCS(runtime.NumCPU())

	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(oauth2.NoContext, ts)

	client := github.NewClient(tc)

	gists, request, _ := client.Gists.List("", &github.GistListOptions{})

	fmt.Println("clone_path: ~/gist-mirror")
	fmt.Println("repositories:")
	printGistInfos(gists)
	for i := 2; i <= request.LastPage; i++ {
		gists, _, _ := client.Gists.List("", &github.GistListOptions{
			ListOptions: github.ListOptions{Page: i},
		})
		printGistInfos(gists)
	}
}

func printGistInfos(gists []*github.Gist) {
	for _, gist := range gists {
		if gist.Description != nil {
			if *gist.Description != "" {
				fmt.Printf("  - url: %v\n    path: %v # %v\n", *gist.GitPullURL, *gist.ID, *gist.Description)
			} else {
				fmt.Printf("  - url: %v\n    path: %v\n", *gist.GitPullURL, *gist.ID)
			}
		} else {
			fmt.Printf("  - url: %v\n    path: %v\n", *gist.GitPullURL, *gist.ID)
		}
	}
}
