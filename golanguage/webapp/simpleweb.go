package main

import (
	"fmt"
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "go is neat")
}
func aboutPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "its all about fun")
}
func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/about/", aboutPage)
	http.ListenAndServe(":8000", nil)
}
