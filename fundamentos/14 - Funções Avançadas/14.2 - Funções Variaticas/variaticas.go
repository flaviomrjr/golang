package main

import "fmt"

func soma(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}

	return total
}

func soma2(numeros ...int) {
	fmt.Println(numeros)
}

func escrever(texto string, numeros ...int) {
	for _, numero := range numeros {
		fmt.Println(texto, numero)
	}
}

func main() {
	totalDaSoma := soma(1, 34, 45, 69)
	fmt.Println(totalDaSoma)

	soma2(2, 4)

	escrever("Olá Mundo", 10, 20)
}
