package main

import (
	"net/http"
	"regexp"
)

type route struct {
	method  string
	regex   *regexp.Regexp
	handler http.HandlerFunc
}

var routes = []route{
	newRoute("GET", "/static/css/output.css", serveCSS),
	newRoute("GET", "/budget/([^/]+)", budget),
}

func newRoute(method, pattern string, handler http.HandlerFunc) route {
	return route{method, regexp.MustCompile(pattern), handler}
}

func serveCSS(w http.ResponseWriter, req *http.Request) {
	handler := http.FileServer(http.FS(css))
	handler.ServeHTTP(w, req)
}

func budget(w http.ResponseWriter, req *http.Request) {
    budgetId := req.Context().Value(ctxKey{}).([]string)[0]
	budget := dbLoadBudget(budgetId)
	html.ExecuteTemplate(w, "budget.html", budget)
}
