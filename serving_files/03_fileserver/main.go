package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir(".")))
	http.HandleFunc("/me", me)

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Fatalln(err)
	}
}

func me(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	io.WriteString(w, `<img src="PHOTO_SATYAKI_SEN.jpg">`)
}
