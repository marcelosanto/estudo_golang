package main

import (
	"time"
)

func main() {
	go escrever("Olá mundo!")
	escrever("programando em go!")

}

func escrever(texto string) {
	for i := 0; i < 5; i++ {
		println(texto)
		time.Sleep(time.Second)
	}
}
