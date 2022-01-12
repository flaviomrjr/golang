package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

type cachorro struct {
	Nome  string `json:"nome`
	Raca  string `json:"raca"`
	Idade uint   `json:"idade"`
}

func main() {
	// STRUCT PARA JSON
	c := cachorro{"Will", "Shihtzu", 3}
	fmt.Println(c)

	cachorroJSON, erro := json.Marshal(c)
	if erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(cachorroJSON)

	fmt.Println(bytes.NewBuffer(cachorroJSON))

	// MAP PARA JSON
	c2 := map[string]string{
		"nome": "Toby",
		"raca": "Poodle",
	}

	cachorroJSON2, erro := json.Marshal(c2)

	if erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(bytes.NewBuffer(cachorroJSON2))
}
