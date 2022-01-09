package main

import "fmt"

func main() {
	fmt.Println("Estruturas de Controle")

	numero := 10

	if numero > 15 {
		fmt.Println("Maior que 15")
	} else {
		fmt.Println("Menor ou igual a 15")
	}

	if outroNumero := numero; outroNumero > 0 { // IF INIT - VARIAVEL outroNumero NÃO PODE SER UTILIZADA EM OUTRO LUGAR FORA DO IF
		fmt.Println("Numero é maior que zero")
	} else if numero < -10 {
		fmt.Println("Numero é menor que - 10")
	} else {
		fmt.Println("entre 0 e -10")
	}
}
