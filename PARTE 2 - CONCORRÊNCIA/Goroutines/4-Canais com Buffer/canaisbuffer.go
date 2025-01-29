package main

import "fmt"

func main() {
	canal := make(chan string, 2) //o numero mostra o tamanho no canal

	canal <- "Olá mundo"
	canal <- "Teste"
	canal <- "Deadlook" //aqui causa o deadlook por conta que estora o tamanho do canal

	mensagem := <-canal
	mensagem2 := <-canal

	fmt.Println(mensagem)
	fmt.Println(mensagem2)

}
