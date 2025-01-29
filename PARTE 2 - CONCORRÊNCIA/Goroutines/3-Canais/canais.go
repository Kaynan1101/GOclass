package main

import (
	"fmt"
	"time"
)

func main() {
	canal := make(chan string)
	go imprimindo("Olá mundo", canal)

	//sintaxe melhorada pra ver se está aberto ou não
	for mensagem := range canal {
		fmt.Println(mensagem)
	}

}

func imprimindo(texto string, canal chan string) {
	for i := 0; i < 3; i++ {
		canal <- texto
		time.Sleep(time.Second)
	}

	close(canal)
}
