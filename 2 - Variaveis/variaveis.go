package main

import "fmt"

func main() {
	/*
		Formas de declarar variaveis
		- o Go nao permite a criação de variaveis sem ouso.
	*/
	var variavel1 string = "Variavel 1"
	variavel2 := "Variavel 2"

	fmt.Println(variavel1)
	fmt.Println(variavel2)

	var (
		variavel3 string = "lalala"
		variavel4 string = "lblblb"
	)

	fmt.Println(variavel3, variavel4)

	variavel5, variavel6 := "variavel5", "variavel6"
	fmt.Println(variavel5, variavel6)

	const constante1 string = "constumava ser assim "
	fmt.Println(constante1)

	variavel5, variavel6 = variavel6, variavel5
	fmt.Println(variavel5, variavel6)
}
