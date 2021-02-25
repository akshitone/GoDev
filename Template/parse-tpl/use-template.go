package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	tpl, err := template.ParseFiles("review/parse-tpl/template.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	htmlFile, err := os.Create("review/parse-tpl/index.html")
	if err != nil {
		log.Println("error creating file", err)
	}
	defer htmlFile.Close()

	err = tpl.Execute(htmlFile, nil)
	if err != nil {
		log.Fatalln(err)
	}
}
