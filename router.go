package main

import (
	"context"
	"log"
	"net/http"
	"strings"
)

var router Router

type Router struct {
	routes []route
}

func (m Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s: %s", r.Method, r.URL.Path)
	var allow []string
	for _, route := range routes {
		matches := route.regex.FindStringSubmatch(r.URL.Path)
		if len(matches) > 0 {
			if r.Method != route.method {
				allow = append(allow, route.method)
				continue
			}
			ctx := context.WithValue(r.Context(), ctxKey{}, matches[1:])
			route.handler(w, r.WithContext(ctx))
			return
		}
	}
	if len(allow) > 0 {
		w.Header().Set("Allow", strings.Join(allow, ", "))
		http.Error(w, "405 method not allowed", http.StatusMethodNotAllowed)
	}
	http.NotFound(w, r)
}

// TODO - Why?
type ctxKey struct{}

func init() {
	router = Router{routes: routes}
}
