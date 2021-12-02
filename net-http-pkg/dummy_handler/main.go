package main

import (
	"fmt"
	"net/http"
)

type satyaki int

func (s satyaki) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	st := "Anything can be written in this method"
	fmt.Println(st)
	fmt.Fprintln(w, st)

}

func main() {
	var sat satyaki
	http.ListenAndServe(":8080", sat)
}
