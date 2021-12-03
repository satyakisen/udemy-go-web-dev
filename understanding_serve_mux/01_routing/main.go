package main

import (
	"fmt"
	"net/http"
)

type dog int

func (dog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/dog":
		w.Header().Set("Content-Type", "text/plain")
		fmt.Fprintln(w, "I am dog")
	case "/cat":
		w.Header().Set("Content-Type", "text/plain")
		fmt.Fprint(w, "I am a cat")
	}

}

func main() {
	var d dog
	http.ListenAndServe(":8080", d)
}
