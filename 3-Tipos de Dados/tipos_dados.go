package main

import (
	"errors"
	"fmt"
)

func main() {
	var n1 int32 = 50
	fmt.Println(n1)

	// alias - apelidos
	//int32 = RUNE
	//uint8 = BYTE
	var n2 rune = 8
	var n3 byte = 15
	fmt.Println(n2, n3)

	var n4 float32 = 12.58
	fmt.Println(n4)

	var boleano bool = true //or false (caso vc não atribua valor ele é FALSE)
	fmt.Println(boleano)

	var erro error //retornando <nil>
	fmt.Println(erro)

	var erro2 error = errors.New("Valor no Erro") //atribuindo valor ao erro com o pacote *errors*
	fmt.Println(erro2)

}
