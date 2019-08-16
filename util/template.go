package util

import (
	"html/template"
	"net/http"
)

var template_location = "templates/"

var templates = template.Must(template.ParseFiles(
	template_location+"edit.html",
	template_location+"view.html",
	template_location+"dashboard.html"))

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
