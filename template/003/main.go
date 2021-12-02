package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {

	err := tpl.ExecuteTemplate(os.Stdout, "01.gohtml", nil)

	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "02.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}

}
