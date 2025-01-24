package main

import "fmt"

func fibonacci(posicao uint) uint {
	if posicao <= 1 {
		return posicao
	}

	return fibonacci(posicao-2) + fibonacci(posicao-1)
}

//Fibonacci pega os dois ultimos numeros é faz a soma para o proximo
//ex: 1 1 2 3 5 8 13 ....

func main() {

	posicao := uint(10)
	//trazendo a soma pela posição
	fmt.Println(fibonacci(posicao))

	//percorrendo e mostrando cada valor na sua posição
	for i := uint(0); i <= posicao; i++ {
		fmt.Println(fibonacci(i))
	}

}
