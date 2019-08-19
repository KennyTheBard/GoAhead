package main

import (
	"log"
	"net/http"

	util "./util"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", util.DashboardHandle)
	mux.HandleFunc("/view/", util.MakeHandle(util.ViewHandle))
	mux.HandleFunc("/edit/", util.MakeHandle(util.EditHandle))
	mux.HandleFunc("/save/", util.MakeHandle(util.SaveHandle))
	mux.Handle("/css/", http.StripPrefix("/css", http.FileServer(http.Dir("css/"))))

	log.Fatal(http.ListenAndServe(":8080", mux))
}
