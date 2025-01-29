package main

import (
	"fmt"
	"introducao-testes/enderecos"
)

func main() {
	tipoEndereco := enderecos.TipoDeEnderecos("a Rui Barbosa")
	fmt.Println(tipoEndereco)
}
