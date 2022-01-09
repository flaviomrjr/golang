package main

import "fmt"

func main() {
	fmt.Println("Maps")

	usuario := map[string]string{

		"nome":      "Pedro",
		"sobrenome": "Silva",
	}

	fmt.Println(usuario)
	fmt.Println(usuario["nome"])

	usuario2 := map[string]map[string]string{
		"nome": {
			"primeiro": "João",
			"segundo":  "Silva",
		},
		"curso": {
			"nome": "Engenharia",
		},
	}

	fmt.Println(usuario2)
	delete(usuario2, "nome")
	fmt.Println(usuario2)

	usuario2["signo"] = map[string]string{
		"nome": "Leão",
	}
	fmt.Println(usuario2)
}
