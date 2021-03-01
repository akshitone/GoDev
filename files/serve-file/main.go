package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", imageServe)
	http.ListenAndServe(":8080", nil)
}

func imageServe(res http.ResponseWriter, req *http.Request) {
	http.ServeFile(res, req, "files/assets/image.jpg")
}
