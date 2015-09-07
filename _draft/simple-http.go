package main

import (
  "fmt"
  "io/ioutil"
  "net/http"
  "encoding/json"
)

type JsonData struct {
  Name string `json:"name"`
  Struct struct {
    Integer int `json:"integer"`
  } `json:"struce"`
}

func main() {
  http.HandleFunc("/", hello)
  http.ListenAndServe(":8001", nil)
}

func hello(writer http.ResponseWriter, reader *http.Request) {
  file, err := ioutil.ReadFile("simple-http.json")
  if err != nil {
    fmt.Println(err)
  }

  var data JsonData

  decoder := json.NewDecoder(file)

  decoded := decoder.Decode(&data)

  writer.Write(decoded)
}
