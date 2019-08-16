package main

import (
	"log"
	"net/http"

	util "./util"
)

func main() {
	http.HandleFunc("/view/", util.MakeHandle(util.ViewHandle))
	http.HandleFunc("/edit/", util.MakeHandle(util.EditHandle))
	http.HandleFunc("/save/", util.MakeHandle(util.SaveHandle))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
