package main

import (
	"fmt"
	"time"
)

func main() {
	i := 0

	for i < 10 {
		i++
		fmt.Println("Incrementando 1")
		time.Sleep(time.Second)
	}

	fmt.Println(i)

	for j := 0; j < 10; j += 2 {
		fmt.Println("Incrementando 2", j)
		time.Sleep(time.Second)
	}

	nomes := [3]string{"Flavio", "Paulo", "AB"}

	for indice, nome := range nomes {
		fmt.Println(indice, nome)
	}

	for _, nome := range nomes {
		fmt.Println(nome)
	}

	for indice, letra := range "PALAVRA" {
		fmt.Println(indice, letra)
		fmt.Println(indice, string(letra))
	}

	usuario := map[string]string{
		"nome":      "Leonardo",
		"sobrenome": "Silva",
	}

	for chave, valor := range usuario {
		fmt.Println(chave, valor)
	}

	//LOOP INFINITO
	for {
		fmt.Println("Executando infinitamente")
		time.Sleep(time.Second)
	}
}
