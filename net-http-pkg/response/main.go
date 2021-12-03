package main

import (
	"fmt"
	"log"
	"net/http"
)

type server int

func (server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>Hello World</h1>")
}

func main() {
	var s server

	err := http.ListenAndServe(":8080", s)

	if err != nil {
		log.Fatalln(err)
	}
}
