package main

import (
	"html/template"
	"log"
	"net/http"
)

type server int

var tpl *template.Template

func (server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}
	tpl.Execute(w, r.Form)
}

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func main() {
	var s server
	err := http.ListenAndServe(":8082", s)

	if err != nil {
		log.Fatalln(err)
	}
}
