package main

import "fmt"

func somar(x int8, y int8) int8 {
	return x + y
}

func calculosMatematicos(n1, n2 int8) (int8, int8) {
	soma := n1 + n2
	subtracao := n1 - n2
	return soma, subtracao
}

func main() {
	soma := somar(10, 20)
	fmt.Println(soma)

	var f = func(txt string) string {
		fmt.Println(txt)
		return txt
	}

	resultado := f("Texto da funcao 1 ")
	fmt.Println(resultado)

	// _ -> serve para ignorar alguma variavel
	resultadoSoma, _ := calculosMatematicos(20, 50)
	fmt.Println(resultadoSoma)
}
