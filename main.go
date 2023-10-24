package main

import (
	"embed"
	"fmt"
	"html/template"
    "io/fs"
	"log"
	"net/http"
)

var (
    //go:embed all:templates/*
    templatesFS embed.FS

    //the parsed templates
    html *template.Template
)

const port = 8080

func hello(w http.ResponseWriter, req *http.Request) {
    fmt.Fprint(w, "hello\n")
}

func budget(w http.ResponseWriter, req *http.Request) {
    html.ExecuteTemplate(w, "budget.html", data)
}

func parseTemplates (templates fs.FS) (*template.Template) {
    parsed := template.Must(template.New("").ParseFS(templates, "templates/*.html"))
    return parsed
}

func main() {
    html = parseTemplates(templatesFS)
    http.HandleFunc("/hello", hello)
    http.HandleFunc("/budget", budget)
    log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
