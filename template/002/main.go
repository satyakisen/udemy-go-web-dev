package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	tpl, err := template.ParseFiles("01.gohtml", "02.gohtml")

	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "01.gohtml", nil)

	if err != nil {
		log.Fatalln(err)

	}

	err = tpl.ExecuteTemplate(os.Stdout, "02.gohtml", nil)

	if err != nil {
		log.Fatalln(err)
	}
}
