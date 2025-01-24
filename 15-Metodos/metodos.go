package main

import "fmt"

type usuario struct {
	nome  string
	idade uint
}

// metodo padrão
func (u usuario) salvar() {
	fmt.Printf("Salvando o nome do usuario como %s", u.nome)
}

//metodo com condição
func (u usuario) maiorDeIdade() bool {
	return u.idade >= 18
}

//metodo utilizando ponteiro para alterar um valor
func (u *usuario) niver() {
	u.idade++
}

func main() {
	usuario1 := usuario{"Kaynan", 24}
	fmt.Println(usuario1)

	idade := usuario1.maiorDeIdade()
	fmt.Println(idade)

	usuario1.niver()
	fmt.Println(usuario1.idade)

}
