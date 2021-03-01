package main

import (
	"io"
	"net/http"
)

func dogFunc(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Heyy, It's a dog!")
}

func catFunc(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Heyy, It's a cat!")
}

func main() {
	http.HandleFunc("/dog/", dogFunc)
	http.HandleFunc("/cat", catFunc)
	// http.Handle("/cat", http.HandleFunc(catFunc))

	http.ListenAndServe(":8080", nil)
}
