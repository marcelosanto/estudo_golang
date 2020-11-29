package main

import "fmt"

type usuario struct {
	nome  string
	idade uint8
}

func (u usuario) salvar() {
	fmt.Printf("Salvandoos dados do Usuário %s no banco de dados\n", u.nome)
}

func (u usuario) maiorDeIdade() bool {
	return u.idade >= 18
}

func (u *usuario) fazerAniversario() {
	u.idade++
}

func main() {
	usuario1 := usuario{"Marcelo", 29}
	fmt.Println(usuario1)
	usuario1.salvar()

	maiorDeIdade := usuario1.maiorDeIdade()
	fmt.Println(maiorDeIdade)

	usuario1.fazerAniversario()
	fmt.Println(usuario1.idade)
}
