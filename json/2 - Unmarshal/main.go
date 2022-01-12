package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type cachorro struct {
	Nome  string `json:"nome` // - PARA NÃO EXIBIR O CAMPO NA REQUISIÇÃO
	Raca  string `json:"raca"`
	Idade uint   `json:"idade"`
}

func main() {
	// JSON PARA STRUCT
	cachorroJSON := `{"Nome":"Will","raca":"Shihtzu","idade":3}`

	var c cachorro

	if erro := json.Unmarshal([]byte(cachorroJSON), &c); erro != nil {
		log.Fatal(erro)
	}
	fmt.Println(c)

	// JSON PARA MAP
	cachorroJSON2 := `{"Nome":"Izzy","raca":"Shihtzu","idade":"1"}`

	c2 := make(map[string]string)
	if erro := json.Unmarshal([]byte(cachorroJSON2), &c2); erro != nil {
		log.Fatal(erro)
	}
	fmt.Println(c2)
}
