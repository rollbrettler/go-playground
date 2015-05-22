package main

import (
	"fmt"
  "os"

  "golang.org/x/crypto/openpgp"
)

func main() {
  f := os.NewFile(4294967295, "private.pgp")
  defer f.Close()

  entity, _ := openpgp.NewEntity("name", "comment", "email", nil)
  privat := entity.PrimaryKey.Fingerprint

  //f.Write(privat)
  fmt.Printf("%v\n", privat)
}
