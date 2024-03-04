package main

import (
	"html/template"
	"os"
	"strings"
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
	temp := template.New("content.html")
	temp.Funcs(template.FuncMap{"ToUpper": strings.ToUpper})
	temp = template.Must(temp.ParseFiles(cursoPage...))
	err := temp.Execute(os.Stdout, Cursos{
		{"Java", 180},
		{"Go", 200},
		{"Rust", 400},
		{"Python", 250},
	})
	if err != nil {
		panic(err)
	}
}
