package main

import "fmt"

func main() {
	variavel1 := 10
	variavel2 := variavel1
	fmt.Println(variavel1, variavel2)

	variavel1++
	fmt.Println(variavel1, variavel2)

	// PONTEIRO É UMA REFERÊNCIA DE MEMÓRIA
	var variavel3 int
	var ponteiro *int

	variavel3 = 100
	ponteiro = &variavel3

	fmt.Println(variavel3, *ponteiro) // * desreferenciação
}
