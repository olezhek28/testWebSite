package main

import (
	"fmt"
	"net/http"
)

func mainPage(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Main page")
}

func aboutPage(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "About page!")
}

func handleRequest() {
	http.HandleFunc("/", mainPage)
	http.HandleFunc("/about/", aboutPage)

	http.ListenAndServe(":8080", nil)
}

func main() {
	handleRequest()
}
