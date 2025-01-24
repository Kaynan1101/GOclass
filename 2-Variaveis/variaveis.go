package main

import "fmt"

func main() {
	var variavel1 string = "Carro"
	variavel2 := "Moto"

	fmt.Println(variavel1, variavel2)

	var (
		variavel5 string = "Bike"
		variavel6 string = "Mobilete"
	)

	fmt.Println(variavel5, variavel6)

	variavel7, vavariavel8 := "Ônibus", "Trem"

	fmt.Println(variavel7, vavariavel8)

	const constante1 string = "Constante 1"
	fmt.Println(constante1)

	//para inverter valores basta fazer da seguinte maneira
	variavel1, variavel2 = variavel2, variavel1
	fmt.Println(variavel1, variavel2)
}
