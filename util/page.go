package util

import (
	"io/ioutil"
	"os"
)

var docDir = "./docs"

type Page struct {
	Title string
	Body  []byte
}

func (p *Page) save() error {
	filename := docDir + "/" + p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := docDir + "/" + title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func deletePage(title string) error {
	filename := docDir + "/" + title + ".txt"
	return os.Remove(filename)
}

func renamePage(oldTitle, newTitle string) error {
	oldFilename := docDir + "/" + oldTitle + ".txt"
	newFilename := docDir + "/" + newTitle + ".txt"
	return os.Rename(oldFilename, newFilename)
}
