package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Olá Mundo!"))
}

func main() {
	// ROTAS
	// URL - IDENTIFICADOR DO RECURSO
	// METODO - GET, POST, PUT, DELETE

	http.HandleFunc("/home", home)

	http.HandleFunc("/usuarios", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Carregar pagina de usuários"))
	})

	log.Fatal(http.ListenAndServe(":5000", nil))

}
