package main

import (
	"log"
	"os"
	"strings"
	"text/template"
)

var tpl *template.Template
var fm = template.FuncMap{
	"uc": strings.ToUpper,
	"ft": firstThree,
}

type sage struct {
	Name  string
	Motto string
}

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("tpl.gohtml"))
}

func firstThree(s string) string {
	s = strings.TrimSpace(s)
	return s[:3]
}
func main() {
	b := sage{
		Name:  "Buddha",
		Motto: "The belief of no beliefs",
	}

	g := sage{
		Name:  "Gandhi",
		Motto: "Be the change",
	}

	s := []sage{b, g}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", s)

	if err != nil {
		log.Fatalln(err)
	}
}
