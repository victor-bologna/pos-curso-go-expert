package main

import (
	"html/template"
	"os"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}

type Cursos []Curso

var cursoPage = []string{
	"header.html",
	"content.html",
	"footer.html",
}

func main() {
	template := template.Must(template.New("content.html").ParseFiles(cursoPage...))
	err := template.Execute(os.Stdout, Cursos{
		{"Java", 180},
		{"Go", 200},
		{"Rust", 400},
		{"Python", 250},
	})
	if err != nil {
		panic(err)
	}
}
