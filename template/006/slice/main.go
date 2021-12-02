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
	s := map[int]string{
		1: "Satyaki",
		2: "Ankita",
		3: "Mimi",
		4: "Rimi",
		5: "Priyam",
	}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", s)

	if err != nil {
		log.Fatalln(err)
	}
}
