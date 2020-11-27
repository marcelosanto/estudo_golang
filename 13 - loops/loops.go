package main

import (
	"fmt"
	"time"
)

func main() {
	// i := 0

	// for i < 10 {
	// 	i++
	// 	fmt.Println("Incrementando i", i)
	// 	time.Sleep(time.Second)
	// }
	// fmt.Println(i)

	for j := 0; j < 10; j += 5 {
		fmt.Println("Incementando o J", j)
		time.Sleep(time.Second)
	}

	nomes := []string{"Marcelo", "Maquerle", "Alice"}

	// primeiro item é o indice, o segundo é o valor
	for _, nome := range nomes {
		fmt.Println(nome)
	}

	for indice, letra := range "PALAVRA" {
		fmt.Println(indice, string(letra))
	}

	usuario := map[string]string{
		"nome":      "Marcelo",
		"sobrenome": "Santos",
	}

	for chave, valor := range usuario {
		fmt.Println(chave, valor)
	}
}
