package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Arrays e Slices")

	// ARRAYS
	var array1 [5]int
	array1[0] = 1
	fmt.Println(array1)

	array2 := [5]string{"p1", "p2", "p3", "p4"}
	fmt.Println(array2)

	array3 := [...]int{1, 2, 3} // TAMANHO BASEADO NA QNT DE VALORES
	fmt.Println(array3)

	// SLICES

	slice := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(slice)

	fmt.Println(reflect.TypeOf(slice))
	fmt.Println(reflect.TypeOf(array1))

	slice = append(slice, 18)
	fmt.Println(slice)

	slice2 := array2[1:3]
	fmt.Println(slice2)

	array2[1] = "Alterado"
	fmt.Println(slice2)
}
