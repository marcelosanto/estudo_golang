package main

import "fmt"

type usuario struct {
	nome     string
	idade    uint8
	endereco endereco
}

type endereco struct {
	logradouro string
	numero     uint8
}

func main() {
	fmt.Println("Arquivo structs")

	var u usuario
	u.nome = "Marcello"
	u.idade = 29
	fmt.Println(u)

	enderecoExemplo := endereco{"Travessa das catatuas", 25}

	usuario2 := usuario{"Marcello", 29, enderecoExemplo}
	fmt.Println(usuario2)

	usuario3 := usuario{idade: 29}
	fmt.Println(usuario3)

}
