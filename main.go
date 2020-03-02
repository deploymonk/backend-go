package main

import (
  "fmt"
  "net/http"
  "github.com/gorilla/mux"
)

func homePage(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Welcome to the HomePage!")
  fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
  myRouter := mux.NewRouter().StrictSlash(true)
  myRouter.HandleFunc("/", homePage)
  http.ListenAndServe(":4000", myRouter)
}

func main() {
  handleRequests()
}
