package main

import "fmt"

func funcao1() {
	fmt.Println("Função 1")
}

func funcao2() {
	fmt.Println("Função 2")
}

func mediaAluno(n1, n2 float32) bool {
	fmt.Println("Média calculada. O resultado será apresentado")
	fmt.Println("Estamos entrando na função")

	media := (n1 + n2) / 2

	if media >= 6 {
		return true
	}

	return false
}

func main() {

	defer funcao1()
	funcao2()

	fmt.Println(mediaAluno(5, 7))

}
