// go build pinetry.go
// encfs --extpass=./pinentry $(pwd)/encfs-test/enc $(pwd)/encfs-test/dec
// curl -X POST -d '{"password": "Start"}' http://localhost:8001

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type password struct {
	Password string
}

func main() {
	http.HandleFunc("/", pinentry)
	http.ListenAndServe(":8001", nil)
}

func pinentry(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":
	case "POST":
		handlePost(w, r)
	}
}

func handlePost(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()
	var t password
	err := decoder.Decode(&t)
	if err != nil {
		w.Write([]byte("{'error':'cannot parse json'}"))
	}

	w.Write([]byte("{}"))
	fmt.Println(t.Password)

	os.Exit(0)
}
