package main

import "fmt"

func main() {
	// ARITMETICOS
	soma := 1 + 2
	subtracao := 1 - 2
	divisao := 10 / 4
	multiplicacao := 10 * 5
	restoDaDivisao := 10 % 2

	fmt.Println(soma, subtracao, divisao, multiplicacao, restoDaDivisao)

	var numero1 int16 = 10
	var numero2 int16 = 25
	soma2 := numero1 + numero2
	fmt.Println(soma2)
	// FIM ARITMETICOS

	// ATRIBUIÇÃO

	var variavel1 string = "String"
	variavel2 := "String2"
	fmt.Println(variavel1, variavel2)
	// FIM DOS OPERADORES DE ATRIBUICAO

	// OPERADORES RELACIONAIS
	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 < 2)
	fmt.Println(1 != 2)
	// FIM DOS RELACIONAIS

	// OPERADORES LOGICOS
	fmt.Println("-----------------")
	verdadeiro, falso := true, false
	fmt.Println(verdadeiro && falso)
	fmt.Println(verdadeiro || falso)
	fmt.Println(!verdadeiro)
	// FIM DOS OPERADORES LOGICOS

	// OPERADORES UNÁRIOS
	numero := 10
	numero++
	numero += 15
	fmt.Println(numero)

	numero--
	numero -= 5
	fmt.Println(numero)

	numero *= 3
	fmt.Println(numero)
	// FIM DOS OPERADORES UNÁRIOS

	// OPERADORES TERNARIOS
	var texto string
	if numero > 5 {
		texto = "Maior que 5"
	} else {
		texto = "Menor que 5"
	}
	fmt.Println(texto)
}
