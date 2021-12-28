package main

import "fmt"

type usuario struct {
	nome     string
	idade    uint8
	endereco endereco //struct dentro de struct
}

type endereco struct {
	logradouro string
	numero     uint8
}

func main() {
	fmt.Println("Arquivo Structs")

	var usuario1 usuario
	usuario1.nome = "Flavio"
	usuario1.idade = 27
	fmt.Println(usuario1)

	endereco1 := endereco{"Rua xxx", 0001}

	usuario2 := usuario{"Will", 2, endereco1}
	fmt.Println(usuario2)

	usuario3 := usuario{nome: "Railma"}
	fmt.Println(usuario3)
}
