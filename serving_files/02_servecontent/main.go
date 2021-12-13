package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func me(w http.ResponseWriter, r *http.Request) {
	file, err := os.Open("PHOTO_SATYAKI_SEN.jpg")

	if err != nil {
		http.Error(w, "file not found", 404)
	}

	defer file.Close()

	fs, err := file.Stat()

	if err != nil {
		http.Error(w, "error occured", 404)
	}

	http.ServeContent(w, r, fs.Name(), fs.ModTime(), file)
}

func mypic(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	io.WriteString(w, `<img src="/me.jpg">`)
}

func main() {
	http.HandleFunc("/", mypic)
	http.HandleFunc("/me.jpg", me)

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Fatalln(err)
	}
}
