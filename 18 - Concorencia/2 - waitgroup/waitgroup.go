package main

import (
	"sync"
	"time"
)

func main() {
	var waitGroup sync.WaitGroup
	waitGroup.Add(2)

	go func()  {
		escrever("Olá mundo!")
		waitGroup.Done()
	}()

	go func()  {
		escrever("programando em go!")
		waitGroup.Done()
	}()
	
	waitGroup.Wait()
}

func escrever(texto string) {
	for i := 0; i < 5; i++ {
		println(texto)
		time.Sleep(time.Second)
	}
}
