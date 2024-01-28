package main

func main() {
	canal := make(chan string, 2)

	canal <- "Olá kakaroto"

	mensagem := <- canal

	println(mensagem)
}