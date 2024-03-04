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

func main() {
	template := template.Must(template.New("template.html").ParseFiles("template.html"))
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
