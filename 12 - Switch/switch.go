package main

import "fmt"

func diaDaSemana(numero int) string {
	switch numero {
	case 1:
		return "Domingo"
	case 2:
		return "Segunda-Feira"
	default:
		return "Dia invalido"
	}
}

func diaDaSemana2(numero int) string {
	switch {
	case numero == 1:
		return "Domingo"
	case numero == 2:
		return "Segunda-Feira"
	default:
		return "Dia Invalido"
	}
}

func diaDaSemana3(numero int) string {
	var diaDaSemana string

	switch {
	case numero == 1:
		diaDaSemana = "Domingo"
		//fallthrough pouco utilizado
	case numero == 2:
		diaDaSemana = "Segunda-Feira"
	default:
		diaDaSemana = "Dia Invalido"
	}

	return diaDaSemana
}

func main() {
	fmt.Println("Switch")

	dia := diaDaSemana(2)
	fmt.Println(dia)

	dia2 := diaDaSemana2(3)
	fmt.Println(dia2)

	dia3 := diaDaSemana3(1)
	fmt.Println(dia3)
}
