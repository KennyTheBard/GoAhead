package util

import (
	"io/ioutil"
	"net/http"
	"path/filepath"
	"strings"
)

func DashboardHandle(w http.ResponseWriter, r *http.Request) {
	files, _ := ioutil.ReadDir(docDir)

	filenames := make([]string, len(files))
	for i, f := range files {
		filenames[i] = strings.TrimSuffix(f.Name(), filepath.Ext(f.Name()))
	}

	err := templates.ExecuteTemplate(w, "dashboard.html", filenames)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
