package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type JsonData struct {
	Name   string `json:"name"`
	Struct struct {
		Integer int `json:"integer"`
	} `json:"struce"`
}

func main() {
	http.HandleFunc("/", hello)
	http.ListenAndServe(":8001", nil)
}

func hello(w http.ResponseWriter, r *http.Request) {
	file, err := ioutil.ReadFile("simple-http.json")
	if err != nil {
		fmt.Println(err)
	}

	var data JsonData

	json.Unmarshal(file, &data)

	w.Write(data)
}
