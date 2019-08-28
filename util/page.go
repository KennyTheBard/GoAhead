package util

import (
	"io/ioutil"
	"os"
)

var doc_dir = "./docs"

type Page struct {
	Title string
	Body  []byte
}

func (p *Page) save() error {
	filename := doc_dir + "/" + p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := doc_dir + "/" + title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func deletePage(title string) error {
	filename := doc_dir + "/" + title + ".txt"
	return os.Remove(filename)
}
