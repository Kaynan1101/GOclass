package main

import "fmt"

func inverterNumeros(numero int) int {
	return numero * -1

}

// * significa que é ponteiro
func inverterNumerosPonteiros(numero *int) {
	*numero = *numero * -1
}

func main() {

	//nesse caso criou uma copia copia do numero passado em outra variavel
	n := 20
	ninvertido := inverterNumeros(n)
	fmt.Println(n)
	fmt.Println(ninvertido)

	//aqui foi trocado o valor da variavel apenas utilzando o ponteiro que mostra o end de memoria dela
	n1 := 40
	fmt.Println(n1)
	inverterNumerosPonteiros(&n1)
	fmt.Println(n1)

}
