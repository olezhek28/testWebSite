package main

import (
	"fmt"
	"net/http"
)

type user struct {
	name string
	age uint16
	balance int64
}

func (u *user) getUserInfo() string {
	return fmt.Sprintf("User name is: %s. He is %d and he has %d$.", u.name, u.age, u.balance)
}

func mainPage(w http.ResponseWriter, req *http.Request) {
	testUser := user{name: "Oleg", age: 26, balance: 10000}

	fmt.Fprintf(w, testUser.getUserInfo())
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
