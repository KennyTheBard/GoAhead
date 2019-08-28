package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	util "./util"
)

func main() {
	var port string
	if len(os.Args) > 1 {
		port = os.Args[1]
	} else {
		port = "8080"
	}

	if num, err := strconv.Atoi(port); err != nil || num < 0 || num > 65535 {
		fmt.Println("Port must be a number between 0 and 65535!")
		return
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/", util.DashboardHandle)
	mux.HandleFunc("/view/", util.MakeHandle(util.ViewHandle))
	mux.HandleFunc("/edit/", util.MakeHandle(util.EditHandle))
	mux.HandleFunc("/save/", util.MakeHandle(util.SaveHandle))
	mux.Handle("/css/", http.StripPrefix("/css", http.FileServer(http.Dir("css/"))))
	mux.Handle("/asset/", http.StripPrefix("/asset", http.FileServer(http.Dir("asset/"))))

	fmt.Println("Starting server on port " + port)
	log.Fatal(http.ListenAndServe(":"+port, mux))
}
