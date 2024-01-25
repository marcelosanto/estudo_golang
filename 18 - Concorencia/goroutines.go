package main

import "time"

func main() {
	go escrever("Olá mundo!")
	escrever("programando em go!")
}

func escrever(texto string) {
	for {
		println(texto)
		time.Sleep(time.Second)
	}
}
