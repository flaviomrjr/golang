package main

import (
	"crud/servidor"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// CREATE TABLE usuarios( id int auto_increment primary key, nome varchar(50) not null, email varchar(5) not null ) ENGINE=INNODB;

// go get github.com/gorilla/mux

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/usuarios", servidor.CriarUsuario).Methods(http.MethodPost)

	fmt.Println("Escutando porta 5000")
	log.Fatal(http.ListenAndServe(":5000", router))
}
