package main

import (
	"crud/servidor"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// CREATE TABLE usuarios( id int auto_increment primary key, nome varchar(50) not null, email varchar(50) not null ) ENGINE=INNODB;

// go get github.com/gorilla/mux

// CREATE = POST
// READ = GET
// UPDATE = PUT
// DELETE = DELETE

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/usuarios", servidor.CriarUsuario).Methods(http.MethodPost)
	router.HandleFunc("/usuarios", servidor.BuscarUsuarios).Methods(http.MethodGet)
	router.HandleFunc("/usuarios/{id}", servidor.BuscarUsuario).Methods(http.MethodGet)
	router.HandleFunc("/usuarios/{id}", servidor.AtualizarUsuario).Methods(http.MethodPut)
	router.HandleFunc("/usuarios/{id}", servidor.DeletarUsuario).Methods(http.MethodDelete)

	fmt.Println("Escutando porta 5000")
	log.Fatal(http.ListenAndServe(":5000", router))
}
