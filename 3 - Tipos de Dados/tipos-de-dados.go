package main

import (
	"errors"
	"fmt"
)

func main() {
	numero := 10000000
	fmt.Println(numero)

	var numero2 uint = 1000
	fmt.Println(numero2)

	// alias
	// INT32 = RUNE
	var numero3 rune = 12456
	fmt.Println(numero3)

	// alias
	// INT8 = BYTE
	var numero4 byte = 123
	fmt.Println(numero4)

	// NUMEROS REAIS

	var numeroReal1 float32 = 123.45
	fmt.Println(numeroReal1)

	var numeroReal2 float64 = 123323.45
	fmt.Println(numeroReal2)

	//FIM DOS NUMEROS REAIS

	// STRINGS
	var str string = "Texto"
	fmt.Println(str)

	str2 := "Texto2"
	fmt.Println(str2)

	char := 'A'
	fmt.Println(char)

	// FIM STRINGS

	var texto string
	fmt.Println(texto)

	var booleanol bool
	fmt.Println(booleanol)

	var erro error = errors.New("Erro interno")
	fmt.Println(erro)

}
