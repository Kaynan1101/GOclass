package main

import "fmt"

//isso é o retorno nomeado
func calculosMat(n1 int, n2 int) (soma int, subtracao int) {
	soma = n1 + n2
	subtracao = n1 - n2
	return
}

func main() {

	a, b := calculosMat(10, 20)
	fmt.Println(a, b)
}
