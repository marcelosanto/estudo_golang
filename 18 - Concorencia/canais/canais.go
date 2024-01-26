package main

import (
	"fmt"
	"time"
)

func main() {
	canal := make(chan string)

	go escrever("Oie, eu sou o Goku!!", canal)

	for {
		mensagem, aberto := <-canal

		if !aberto {
			break
		}

		println(mensagem)
	}

	fmt.Println("Fim do programa")

}

func escrever(texto string, canal chan string) {
	for i := 0; i < 5; i++ {
		canal <- texto
		time.Sleep(time.Second)
	}

	close(canal)
}
