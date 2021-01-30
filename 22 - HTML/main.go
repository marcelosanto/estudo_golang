package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type usuario struct {
	Nome  string
	Email string
}

var templates *template.Template

func main() {
	templates = template.Must(template.ParseGlob("*.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		u := usuario{
			"João",
			"joao.pedro@gmail.com",
		}
		templates.ExecuteTemplate(w, "index.html", u)
	})

	fmt.Println("Escutando na porta 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
