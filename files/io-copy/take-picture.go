package main

import (
	"io"
	"net/http"
	"os"
)

func fileFunc(res http.ResponseWriter, req *http.Request) {
	file, err := os.Open("files/assets/image.jpg")
	if err != nil {
		http.Error(res, "File not found!", 404)
		return
	}
	defer file.Close()

	io.Copy(res, file)
}
func filePathFunc(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(res, `
	<img src="files/assets/image.jpg" />
	`)
}

func main() {
	http.HandleFunc("/", fileFunc)
	http.HandleFunc("/img", filePathFunc)
	http.ListenAndServe(":8080", nil)
}
