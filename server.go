package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)

type Page struct {
	Title string
	Body  []byte
}

func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func main() {
	http.Handle("/", http.FileServer(http.Dir("./webpage")))
	http.HandleFunc("/view/", viewHandle)
	http.HandleFunc("/edit/", editHandle)
	// http.HandleFunc("/save/", saveHandle)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func viewHandle(w http.ResponseWriter, t *http.Request) {
	title := t.URL.Path[len("/view/"):]
	p, _ := loadPage(title)

	fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
}

func editHandle(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/edit/"):]
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}

	t, err := template.ParseFiles("webpage/edit.html")
	if err != nil {
		fmt.Println(err)
	} else {
		t.Execute(w, p)
	}
}

// func saveHandle(w http.ResponseWriter, t *http.Request) {
// 	title := t.URL.Path[len("/view/"):]
// 	p, _ := loadPage(title)

// 	fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
// }
