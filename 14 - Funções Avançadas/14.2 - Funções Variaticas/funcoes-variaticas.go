package main

import "fmt"

func soma(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}

	return total
}

// parametro variatico deve ser o ultimo argumento da funcao
func escrever(texto string, numeros ...int) {
	for _, numero := range numeros {
		fmt.Println(texto, numero)
	}
}

func main() {
	totalDaSoma := soma(1, 2, 3, 5, 9, 88, 77, 66, 33)
	fmt.Println(totalDaSoma)

	escrever("Olá mundo!", 55, 11, 42, 44, 78)

}
