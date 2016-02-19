// go build pinetry.go
// encfs --extpass=./pinentry $(pwd)/encfs-test/enc $(pwd)/encfs-test/dec
// curl -X POST -d '{"password": "Start"}' http://localhost:8001

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"text/template"
)

var htmlTemplate = `
<!DOCTYPE html>
<html>
  <head>
    <meta charset="utf-8">
    <title>Pinentry</title>
  </head>
  <body>
    <form action="/" method="post">
			<input name="password" type="password" />
			<button>Submit pin</button>
		</form>
  </body>
</html>
`

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
		t, _ := template.New("html").Parse(htmlTemplate)
		t.Execute(w, nil)
	case "POST":
		handlePost(w, r)
	}
}

func handlePost(w http.ResponseWriter, r *http.Request) {

	// If the form was send via http post use this
	if r.PostFormValue("password") != "" {
		fmt.Println(r.PostFormValue("password"))
		w.Write([]byte("password was submitted"))
		os.Exit(0)
	}

	// Try to fetch the password paramerter from the json body
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()
	var t password
	err := decoder.Decode(&t)
	if err != nil {
		w.Write([]byte("{'error':'cannot parse json'}"))
		return
	}

	w.Write([]byte("{'message':'password was submitted'}"))
	fmt.Println(t.Password)

	os.Exit(0)
}
