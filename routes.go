package main

import (
	"net/http"
	"regexp"
)

type route struct {
    method string
    regex *regexp.Regexp
    handler http.HandlerFunc
}

var routes = []route{
    newRoute("GET", "/buget/([^/]+)", budget),
}

func newRoute(method, pattern string, handler http.HandlerFunc) route {
    return route{method, regexp.MustCompile(pattern), handler}
}


func budget(w http.ResponseWriter, req *http.Request) {
    html.ExecuteTemplate(w, "budget.html", data)
}

