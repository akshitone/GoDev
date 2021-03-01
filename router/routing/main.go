package main

import (
	"io"
	"net/http"
)

type onetype int

func (one onetype) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	switch request.URL.Path {
	case "/dog":
		io.WriteString(response, "It's a dog!")
	case "/cat":
		io.WriteString(response, "It's a cat!")
	}
}

func main() {
	var a onetype
	http.ListenAndServe(":8080", a)
}
