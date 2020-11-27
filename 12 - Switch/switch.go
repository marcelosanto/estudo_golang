package main

import "fmt"

func diaDaSemana(numero int) string {
	switch numero {
	case 1:
		return "Domingo"
	case 2:
		return "Segunda-Feira"
	case 3:
		return "Terça-Feira"
	case 4:
		return "Quarta-Feira"
	case 5:
		return "Quinta-Feira"
	case 6:
		return "Sexta-Feira"
	case 7:
		return "Sabado"
	default:
		return "Número inválido"
	}
}

func diaDaSemana2(numero int) string {

	var diaSemana string

	switch {
	case numero == 1:
		diaSemana = "Domingo"
	case numero == 2:
		diaSemana = "Segunda-Feira"
	case numero == 3:
		diaSemana = "Terça-Feira"
	case numero == 4:
		diaSemana = "Quarta-Feira"
	case numero == 5:
		diaSemana = "Quinta-Feira"
	case numero == 6:
		diaSemana = "Sexta-Feira"
	case numero == 7:
		diaSemana = "Sabado"
	default:
		diaSemana = "Número inválido"
	}

	return diaSemana
}

func main() {
	dia := diaDaSemana(7)
	fmt.Println(dia)

	fmt.Println("--------")

	dia2 := diaDaSemana2(1)
	fmt.Println(dia2)
}
