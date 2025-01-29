package main

import "fmt"

//função variatica
func soma(numeros ...int) int {
	total := 0
	for _, numeros := range numeros {
		total += numeros
	}

	return total
}

func escrevendo(texto string, numeros ...int) {
	for _, numero := range numeros {
		fmt.Println(texto, numero)
	}
}

func main() {

	tSoma := soma(1, 2, 3, 4)
	fmt.Println(tSoma)

	escrevendo("Olá mundo", 1, 2, 3, 4, 5)

}
