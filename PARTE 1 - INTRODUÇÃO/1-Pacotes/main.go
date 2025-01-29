package main

import (
	"fmt"
	"modulo/auxiliar"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Testando Pacotes e Módulos")
	auxiliar.Escrevendo()

	erro := checkmail.ValidateFormat("kaynan.lima1101@hotmail.com")
	fmt.Println(erro)
}
