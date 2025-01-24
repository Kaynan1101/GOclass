package main

import "fmt"

func main() {

	//função que não tem nome
	retorno := func(texto string) string {
		return fmt.Sprintf("Recebido -> %s", texto)
	}("Passando Paramêtro")

	fmt.Println(retorno)
}
