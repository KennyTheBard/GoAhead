package main

import (
	"fmt"
	"net/http"

	html "./html"
)

func main() {
	http.HandleFunc("/", handle)

	fmt.Println("Server is on!")

	http.ListenAndServe(":21337", nil)
}

func handle(w http.ResponseWriter, t *http.Request) {

	content := html.CreateContent("Hello Bucureeesti")
	para := html.CreateContent("Ce faceti cum mai sunteti?")
	content.Tag("h1").Merge(para.Tag("p")).Tag("body").Tag("html")

	fmt.Fprintln(w, content.Text)
}
