package main

import (
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", imageServe)
	http.ListenAndServe(":8080", nil)
}

func imageServe(res http.ResponseWriter, req *http.Request) {
	file, err := os.Open("files/assets/image.jpg")
	if err != nil {
		http.Error(res, "Image not found!", 404)
		return
	}
	defer file.Close()

	fileState, err := file.Stat()
	if err != nil {
		http.Error(res, "Image not found!", 404)
	}

	http.ServeContent(res, req, file.Name(), fileState.ModTime(), file)
}
