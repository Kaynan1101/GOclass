package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	//recebendo a entrada de dois canais através da função multiplexar e dando a saida em apenas um
	canal := multiplexar(escrever("Olá mundo"), escrever("Testando"))

	for i := 0; i < 10; i++ {
		fmt.Println(<-canal)
	}

}

// dois canais de entrada que envia para apenas 1
func multiplexar(canaDeEntrada1, canalDeEntrada2 <-chan string) <-chan string {
	canalDeSaida := make(chan string)

	go func() {
		for {
			select {
			case mensagem := <-canaDeEntrada1:
				canalDeSaida <- mensagem
			case mensagem := <-canalDeEntrada2:
				canalDeSaida <- mensagem
			}
		}
	}()

	return canalDeSaida
}

func escrever(texto string) <-chan string {
	canal := make(chan string)

	go func() {
		for {
			canal <- fmt.Sprintf("Valor recebido: %s", texto)
			//criando um tempo de duração random na saída
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(2000)))
		}
	}()

	return canal

}
