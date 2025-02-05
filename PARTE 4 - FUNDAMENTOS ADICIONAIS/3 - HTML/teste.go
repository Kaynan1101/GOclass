package main

import (
	"log"
	"net/http"
	"text/template"
)

var templates *template.Template

type usuario struct {
	Nome string
}

func main() {

	//Renderizando um arquivo HTML
	templates = template.Must(template.ParseGlob("*.html"))

	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {

		//Deixando o HTML Dinamico
		u := usuario{"Kaynan"}

		//Execute.Template(w, nome do arquivo html, dinamica)
		templates.ExecuteTemplate(w, "home.html", u)
	})

	log.Fatal(http.ListenAndServe(":5000", nil))

}
