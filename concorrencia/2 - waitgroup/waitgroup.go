package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// CONCORRENCIA != PARALELISMO

	var waitGroup sync.WaitGroup

	waitGroup.Add(2)

	go func() {
		escrever("Olá mundo")
		waitGroup.Done() // -1
	}()

	go func() {
		escrever("Programando em Go!")
		waitGroup.Done() // -1
	}()

	waitGroup.Wait() // Espera as gorouines serem executadas
}

func escrever(texto string) {
	for i := 0; i < 5; i++ {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
