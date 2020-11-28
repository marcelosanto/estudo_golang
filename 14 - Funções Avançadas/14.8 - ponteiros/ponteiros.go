package main

import "fmt"

func inverterSinal(numero int) int {
	return numero * -1
}

// * ponteiro aponta para o endereco de memoria
func inverterSinalComPonteiro(numero *int) {
	*numero *= -1
}

func main() {
	numero := 20
	numeroInvertido := inverterSinal(numero)
	fmt.Println(numeroInvertido)

	novoNumero := 40
	fmt.Println(novoNumero)
	inverterSinalComPonteiro(&novoNumero)
	fmt.Println(novoNumero)
}
