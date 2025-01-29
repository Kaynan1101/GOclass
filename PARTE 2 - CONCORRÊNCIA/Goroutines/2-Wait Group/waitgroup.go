package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	var waitGroup sync.WaitGroup

	waitGroup.Add(2)

	go func() {
		imprimindo("Olá mundo") //go = goroutines
		waitGroup.Done()

	}()
	go func() {
		imprimindo("Concorrência")
		waitGroup.Done()
	}()

	waitGroup.Wait()
}

func imprimindo(texto string) {
	//loop infinito para testar a concorrencia
	for i := 0; i <= 5; i++ {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
