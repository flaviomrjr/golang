package main

import "fmt"

func main() {
	/*func(texto string) {

		fmt.Println(texto)
	}("Passando parametro")*/

	retorno := func(texto2 string) string {
		return fmt.Sprintf("Recebido -> %s", texto2)
	}("Retornos")

	fmt.Println(retorno)
}
