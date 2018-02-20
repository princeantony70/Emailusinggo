package main

import (
    "net/http"
)

func Group(res http.ResponseWriter, req *http.Request) {
    println("group handler")
}

func main() {
    http.HandleFunc("/group/", Group)
    err := http.ListenAndServe(":9001", nil)
    if err != nil {
      panic(err)
    }
    println("Running code after ListenAndServe (only happens when server shuts down)")
}
