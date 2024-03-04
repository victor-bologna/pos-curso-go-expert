package main

import (
	"os"
	"text/template"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}

func main() {
	golang := Curso{"Go", 20}
	template := template.Must(template.New("Curso Template").Parse("Curso: {{.Nome}} - Carga Horária: {{.CargaHoraria}}"))
	err := template.Execute(os.Stdout, golang)
	if err != nil {
		panic(err)
	}
}
