package main

import (
	"fmt"
	"time"
)

func main() {
	i := 0

	//Utilizando Loop padrão
	for i < 3 {
		i++
		fmt.Println("Incrementando I")
		time.Sleep(time.Second) //função de tempo
	}

	fmt.Println(i)

	//Usando Aninhado com if inits
	for j := 0; j < 3; j += 2 {
		fmt.Println("Incrementando J", j)
		time.Sleep(time.Second)
	}

	//Utilizando o Range

	nomes := [2]string{"Kaynan", "Fernanda"}

	for indice, nome := range nomes {
		fmt.Println(indice, nome)
	}

	//caso queira que apareça só o nome sem o indice deve se usar o underline
	for _, nome := range nomes {
		fmt.Println(nome)
	}

	//se imprimir assim vai aparecer as letras no modelo de tabela ask
	for indice, letra := range "KAYNAN" {
		fmt.Println(indice, letra)
	}

	// dessa forma aparece as letras reais
	for indice, letra := range "KAYNAN" {
		fmt.Println(indice, string(letra))
	}

}
