package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type User struct {
	Name string
	Age uint16
	Balance int64
}

func (u *User) getUserInfo() string {
	return fmt.Sprintf("User name is: %s. He is %d and he has %d$.", u.Name, u.Age, u.Balance)
}

func mainPage(w http.ResponseWriter, req *http.Request) {
	testUser := User{Name: "Oleg", Age: 26, Balance: 10000}

	tmpl, _ := template.ParseFiles("templates/mainPage.html")
	tmpl.Execute(w, testUser)
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
