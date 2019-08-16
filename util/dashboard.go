package util

import (
	"io/ioutil"
	"net/http"
	"path/filepath"
	"strings"
)

type file_meta struct {
	Title string
}

type files_wrapper struct {
	Pages []file_meta
}

func DashboardHandle(w http.ResponseWriter, r *http.Request) {
	files, _ := ioutil.ReadDir(doc_dir)

	filenames := make([]file_meta, len(files))
	for i, f := range files {
		filenames[i].Title = strings.TrimSuffix(f.Name(), filepath.Ext(f.Name()))
	}

	err := templates.ExecuteTemplate(w, "dashboard.html", files_wrapper{Pages: filenames})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
