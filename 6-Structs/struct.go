package main

import "fmt"

type usuario struct {
	nome     string
	idade    uint8
	endereco endereco
}

type endereco struct {
	logradouro string
	numero     int
}

func main() {
	fmt.Println("Arquivo de Structs")

	var u usuario
	u.idade = 18
	u.nome = "Kay"
	fmt.Println(u)

	enderecoTeste := endereco{"Rua José Scali", 172}
	fmt.Println(enderecoTeste)

	u2 := usuario{"Fer", 20, enderecoTeste}
	fmt.Println(u2)

	u3 := usuario{nome: "Lima"}
	fmt.Println(u3)

}
