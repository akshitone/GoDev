package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("review/data-tpl/template-structs.gohtml"))
}

type civil struct {
	Name string
	Age  int
}

type nineNine struct {
	civil
	LicenceToKill bool
}

type person struct {
	Civils    []civil
	NineNines []nineNine
}

func main() {
	akshit := civil{"akshit", 21}
	viral := civil{"viral", 22}

	jake := nineNine{civil: civil{"jake", 32}, LicenceToKill: true}
	amy := nineNine{civil: civil{"amy", 31}, LicenceToKill: true}

	person := struct {
		Civils    []civil
		NineNines []nineNine
	}{
		[]civil{akshit, viral},
		[]nineNine{jake, amy},
	}

	err := tpl.Execute(os.Stdout, person)
	if err != nil {
		log.Fatalln(err)
	}
}
