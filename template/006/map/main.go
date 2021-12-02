package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	s := []string{"Satyaki", "Ankita", "Mimi", "Rimi", "Priyam"}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", s)

	if err != nil {
		log.Fatalln(err)
	}
}
