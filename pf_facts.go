package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mphox/pf_facts/handler/root"
	"github.com/mphox/pf_facts/handler/rule"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", root.RootHandler)
	r.HandleFunc("/rule/{int}", rule.RuleHandler)
	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}
