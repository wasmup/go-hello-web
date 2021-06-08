package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", greet)
	http.HandleFunc("/favicon.ico", favicon)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func greet(w http.ResponseWriter, r *http.Request) {
	s := r.URL.Path[1:]
	if len(s) == 0 {
		s = "World"
	}
	io.WriteString(w, "Hello "+s)
}

func favicon(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "favicon.ico")
}
