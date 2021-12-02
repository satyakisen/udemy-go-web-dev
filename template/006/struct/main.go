package main

import (
	"log"
	"os"
	"text/template"
)

type sage struct {
	Name     string
	Location string
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	sage1 := sage{
		"Gandhi",
		"India",
	}

	sage2 := sage{
		"Buddha",
		"Nepal",
	}

	listOfSage := []sage{sage1, sage2}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", listOfSage)

	if err != nil {
		log.Fatalln(err)
	}
}
