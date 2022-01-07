package main

import "fmt"

func closure() func() {
	texto := "Função closure"

	funcao := func() {
		fmt.Println(texto)
	}

	return funcao
}

func main() {
	texto := "Função main"
	fmt.Println(texto)

	funcaoNova := closure()
	funcaoNova()
}
