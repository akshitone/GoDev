package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("func-tpl/tpl-fun.gohtml"))
}

type person struct {
	Name string
	Age  int
}

func (p person) AgeDouble() int {
	return p.Age * 2
}

func (p person) TakesArg(x int) int {
	return x * 2
}

func (p person) SomeProcess() int {
	return 8
}

func main() {
	akshit := person{"akshit", 21}
	err := tpl.Execute(os.Stdout, akshit)
	if err != nil {
		log.Fatalln(err)
	}
}
