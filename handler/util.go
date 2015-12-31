package handler

import (
	"html/template"
	"net/http"

	"github.com/mphox/pf_facts/model"
)

var templates = template.Must(template.ParseFiles("templates/rule.html"))

func RenderTemplate(w http.ResponseWriter, template string, rule *model.Rule) {
	if err := templates.ExecuteTemplate(w, template, rule); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
