package main

import (
	"net/http"
	"text/template"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}

type Cursos []Curso

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		t := template.Must(template.New("template.html").Parse(`O curso {{.Nome}} tem {{.CargaHoraria}} horas de carga hora`))
		err := t.Execute(w, Cursos {
			{"Go", 40},
			{"Java", 20},
			{"Python", 30},
		})
		if err != nil {
			panic(err)
		}
	})
	http.ListenAndServe(":8282", nil)
}