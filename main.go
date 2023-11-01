package main

import (
	"database/sql"
	"embed"
	"fmt"
	"html/template"
	"io/fs"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

var (
	//go:embed all:templates/*
	templatesFS embed.FS

	//go:embed static/css/output.css
	css embed.FS

	//the parsed templates
	html *template.Template

	db *sql.DB
)

const port = 8080
const dbPath = "budgets.db"

func init() {
	var err error
	db, err = sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatal(err)
	}
}

func parseTemplates(templates fs.FS) *template.Template {
	parsed := template.Must(template.New("").ParseFS(templates, "templates/*.html"))
	return parsed
}

func main() {
	html = parseTemplates(templatesFS)

	http.Handle("/static/css/output.css", http.FileServer(http.FS(css)))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), router))

	defer db.Close()
}
