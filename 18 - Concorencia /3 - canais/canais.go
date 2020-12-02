package main

import (
	"fmt"
	"time"
)

func main() {
	canal := make(chan string)

	go escrever("Oieeee", canal)

	fmt.Println("Depois da função escrever comecar a ser executada")

	for {
		msg, aberto := <-canal
		if !aberto {
			break
		}
		fmt.Println(msg)
	}

	fmt.Println("Fim do programa")

}

func escrever(texto string, canal chan string) {
	time.Sleep(time.Second * 5)
	for i := 0; i < 5; i++ {
		canal <- texto
		time.Sleep(time.Second)
	}

	close(canal)
}
