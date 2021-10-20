package main

import (
	"errors"
	"fmt"
)

func main() {
	// NUMEROS INTEIROS
	var numero int64 = 1000000000
	fmt.Println(numero)

	//uint unsygned int

	var numero2 uint32 = 10000
	fmt.Println(numero2)

	// alias
	// int32 = rune
	var numero3 rune = 12344
	fmt.Println(numero3)

	// int8 = byte
	var numero4 byte = 23
	fmt.Println(numero4)

	//FIM NUMEROS INTEIROS

	// NUMEROS REAIS
	var real1 float32 = 123.45
	fmt.Println(real1)

	var real2 float64 = 3456.34
	fmt.Println(real2)

	real3 := 1234.56
	fmt.Println(real3)

	//FIM NUMEROS REAIS

	// STRINGS

	var str string = "Texto"
	fmt.Println(str)

	str2 := "Text2"
	fmt.Println(str2)

	//FIM STRINGS

	var texto string
	fmt.Println(texto)

	//BOOL

	var bool1 bool
	fmt.Println(bool1)

	//ERROR
	var erro error = errors.New("Error Interno")
	fmt.Println(erro)

}
