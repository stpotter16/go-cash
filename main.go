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

    //go:embed static/css/output.css
    css embed.FS

    //the parsed templates
    html *template.Template
)

const port = 8080

func parseTemplates (templates fs.FS) (*template.Template) {
    parsed := template.Must(template.New("").ParseFS(templates, "templates/*.html"))
    return parsed
}

func main() {
    html = parseTemplates(templatesFS)

    http.Handle("/static/css/output.css", http.FileServer(http.FS(css)))
    log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), router))
}
