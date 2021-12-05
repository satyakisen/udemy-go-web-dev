package main

import (
	"fmt"
	"log"
	"net/http"
)

func i(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "This is index page")
}

func d(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hi there! I am dogo...")
}

func m(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "This is Satyaki Sen")
}

func main() {
	http.HandleFunc("/", i)
	http.HandleFunc("/dog/", d)
	http.HandleFunc("/me/", m)
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Fatalln(err)
	}
}
