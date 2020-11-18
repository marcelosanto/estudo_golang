package main

import (
	"fmt"
)

func main() {
	fmt.Println("Arrays e Slices")

	var array1 [5]string
	array1[0] = "Marcelo"
	fmt.Println(array1)

	array2 := [5]string{"Um", "Dois", "Tres", "Quatro", "Cinco"}
	fmt.Println(array2)

	array3 := [...]int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(array3)

	slice := []int{100, 11, 22, 333, 44, 88}
	fmt.Println(slice)

	slice = append(slice, 117)
	fmt.Println(slice)

	slice2 := array2[1:3]
	fmt.Println(slice2)

	// Arrays Internos
	slice3 := make([]float32, 10, 11)
	fmt.Println(slice3)

	slice3 = append(slice3, 88)
	slice3 = append(slice3, 175)

	fmt.Println(slice3)
	fmt.Println(len(slice3)) // -> length
	fmt.Println(cap(slice3)) // -> capacidade

	// quando nao se passa a capacidade, ela e automaticamente definida pelo tamanho do slice
	slice4 := make([]float32, 5)
	fmt.Println(slice4)
	fmt.Println(len(slice4)) // -> length
	fmt.Println(cap(slice4)) // -> capacidade

}
