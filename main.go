package main

import (
    "fmt"
    "log"
    "net/http"
)

const port = 8080

func hello(w http.ResponseWriter, req *http.Request) {
    fmt.Fprint(w, "hello\n")
}

func main() {
    http.HandleFunc("/hello", hello)
    log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
