package util

import (
	"html/template"
	"net/http"
)

var templates = template.Must(template.ParseFiles(
	"templates/head.html",
	"templates/edit.html",
	"templates/view.html",
	"templates/dashboard.html"))

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
