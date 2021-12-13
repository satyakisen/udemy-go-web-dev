package main

import (
	"io"
	"log"
	"net/http"
)

func me(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "PHOTO_SATYAKI_SEN.jpg")
}

func mypic(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	io.WriteString(w, `<img src="/me.jpg">`)
}

func main() {
	http.HandleFunc("/", mypic)
	http.HandleFunc("/me.jpg", me)

	err := http.ListenAndServe(":8082", nil)

	if err != nil {
		log.Fatalln(err)
	}
}
