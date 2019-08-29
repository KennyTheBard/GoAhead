package util

import (
	"net/http"
	"strings"
)

func MakeHandle(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		pathArgs := strings.Split(r.URL.Path, "/")
		if len(pathArgs) <= 1 {
			http.NotFound(w, r)
			return
		}

		fn(w, r, pathArgs[2])
	}
}

func EditHandle(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}

	renderTemplate(w, "edit", p)
}

func ViewHandle(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadPage(title)
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}

	renderTemplate(w, "view", p)
}

func SaveHandle(w http.ResponseWriter, r *http.Request, title string) {
	body := r.FormValue("body")
	p := &Page{Title: title, Body: []byte(body)}
	err := p.save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

func DeleteHandle(w http.ResponseWriter, r *http.Request, title string) {
	deletePage(title)

	http.Redirect(w, r, "/", http.StatusFound)
}

func RenameHandle(w http.ResponseWriter, r *http.Request) {
	if pathArgs := strings.Split(r.URL.Path, "/"); len(pathArgs) == 4 {
		oldTitle := pathArgs[2]
		newTitle := pathArgs[3]

		renamePage(oldTitle, newTitle)
		http.Redirect(w, r, "/view/"+string(newTitle), http.StatusFound)

	} else {
		http.Redirect(w, r, "/", http.StatusNotFound)
	}
}
