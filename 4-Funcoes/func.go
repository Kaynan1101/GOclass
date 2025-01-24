package main

import "fmt"

func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}

func Calculos(n1, n2 int32) (int32, int32) {
	soma := n1 + n2
	sub := n1 - n2
	return soma, sub
}

func main() {
	soma := somar(5, 8)
	fmt.Println(soma)

	var f = func() {
		fmt.Println("Função F")
	}

	f()

	resultadosSoma, _ := Calculos(20, 30)
	fmt.Println(resultadosSoma)

}
