package main

import (
	"html/template"
	"net/http"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}

type Cursos []Curso

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", cursosPage)
	http.ListenAndServe(":8080", mux)
}

func cursosPage(writer http.ResponseWriter, request *http.Request) {
	template := template.Must(template.New("template.html").ParseFiles("template.html"))
	err := template.Execute(writer, Cursos{
		{"Java", 300},
		{"Go", 200},
		{"Rust", 500},
		{"Python", 450},
	})
	if err != nil {
		panic(err)
	}
}
