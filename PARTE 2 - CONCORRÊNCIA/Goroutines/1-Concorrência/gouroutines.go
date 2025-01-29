package main

import (
	"fmt"
	"time"
)

func main() {
	//CONCORRÊNCIA != PARALELISMO

	go imprimindo("Olá mundo") //go = goroutines
	imprimindo("Concorrência")
}

func imprimindo(texto string) {
	//loop infinito para testar a concorrencia
	for {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
