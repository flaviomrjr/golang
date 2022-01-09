package main

import (
	"fmt"
)

type pessoa struct {
	nome      string
	sobrenome string
	idade     uint8
	altura    uint8
}

type estudante struct {
	pessoa
	curso     string
	faculdade string
}

func main() {
	fmt.Println("Herança")

	pessoa1 := pessoa{"Flavio", "Rocha", 47, 177}
	fmt.Println(pessoa1)

	estudante1 := estudante{pessoa1, "Medicina", "USP"}
	fmt.Println(estudante1)
	fmt.Println(estudante1.nome)
}
