package main

import (
	"html/template"
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {

	u := usuario{"João", "joao.pedro@gmail.com"}

	templates.ExecuteTemplate(w, "home.html", u)
}

type usuario struct {
	Nome  string
	Email string
}

var templates *template.Template

func main() {
	templates = template.Must(template.ParseGlob("*.html"))

	http.HandleFunc("/home", home)

	log.Fatal(http.ListenAndServe(":5000", nil))

}
