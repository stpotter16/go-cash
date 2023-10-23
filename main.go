package main

import (
    "fmt"
    "html/template"
    "log"
    "net/http"
)

const port = 8080

func hello(w http.ResponseWriter, req *http.Request) {
    fmt.Fprint(w, "hello\n")
}

func budget(w http.ResponseWriter, req *http.Request) {
    tmpl := template.Must(template.ParseFiles("templates/budget.html"))
    tmpl.Execute(w, nil)
}

func main() {
    http.HandleFunc("/hello", hello)
    http.HandleFunc("/budget", budget)
    log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
