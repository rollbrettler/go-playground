package main

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
)

func main() {
	encfsPath, err := exec.LookPath("encfs")
	if err != nil {
		log.Fatal("Please install encfs")
	}
	cmd := exec.Command(encfsPath, "-v", "--reverse", "--stdinpass", "$ENCFSPASSWORD", "~/mobile", "~/crypt")
	var out bytes.Buffer
	cmd.Stdout = &out
	err = cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("in all caps: %q\n", out.String())
}
