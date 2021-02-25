package main

import (
	"log"
	"math"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.New("").Funcs(funMap).ParseFiles("review/pipeline-tpl/tpl-pipeline.gohtml"))
}

var funMap = template.FuncMap{
	"funDouble":     funDouble,
	"funSquare":     funSquare,
	"funSquareRoot": funSquareRoot,
}

func funDouble(x int) int {
	return x + x
}

func funSquare(x int) float64 {
	return math.Pow(float64(x), 2)
}

func funSquareRoot(x float64) float64 {
	return math.Sqrt(x)
}

func main() {
	err := tpl.ExecuteTemplate(os.Stdout, "tpl-pipeline.gohtml", 100)
	if err != nil {
		log.Fatalln(err)
	}
}
