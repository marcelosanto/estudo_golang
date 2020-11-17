package main

import "fmt"

func main() {
	// ARITMETICOS
	soma := 1 + 2
	subatracao := 1 - 2
	divisao := 10 / 2
	multiplicacao := 10 * 5
	resDaDivisao := 10 % 2

	fmt.Println(soma, subatracao, divisao, multiplicacao, resDaDivisao)

	var numero1 int16 = 10
	var numero2 int16 = 25
	soma2 := numero1 + numero2
	fmt.Println(soma2)

	// FIM DOS ARITMETICOS

	// ATRIBUIÇÃO
	var variavel1 string = "String"
	variavel2 := "String2"
	fmt.Println(variavel1, variavel2)
	// FIM DOS OPERADORS DE ATRIBUICAO

	// OPERADORES RELACIONAIS
	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 < 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 != 2)
	// FIM DOS RELACIONAIS

	// OPERADORES LOGICOS
	verdadeiro, falso := true, false
	fmt.Println(verdadeiro && falso)
	fmt.Println(verdadeiro || falso)
	fmt.Println(!verdadeiro)
	fmt.Println(!falso)
	// FIM DOS OPERADORES LOGICOS

	//OPERADORES UNARIOS
	fmt.Println("-------------------")
	numero := 10
	numero++
	numero += 15
	fmt.Println(numero)

	numero--
	numero -= 20

	numero *= 3
	numero /= 10
	numero %= 3
	fmt.Println(numero)
	// FIM DOS OPERADORES UNARIOS

}
