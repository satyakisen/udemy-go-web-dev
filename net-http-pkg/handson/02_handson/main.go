package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func me(w http.ResponseWriter, r *http.Request) {
	err := tpl.Execute(w, "I am Satyaki Sen")

	if err != nil {
		log.Fatalln(err)
	}
}

func main() {

	http.HandleFunc("/me/", me)

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Fatalln(err)
	}

}
