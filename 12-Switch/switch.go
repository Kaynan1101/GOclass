package main

import "fmt"

func marcaCarros(numero int) string {
	switch numero {
	case 1:
		return "Ferrari"
	case 2:
		return "Lamborghini"
	case 3:
		return "Porche"
	default:
		return "Numero Invalido!"
	}
}

//outra maneira
func marcaCarros2(numero int) string {
	switch {
	case numero == 1:
		return "BMW"

	case numero == 2:
		return "Mercedes"

	default:
		return "Numero Invalido!"
	}
}

func main() {

	carro := marcaCarros(1)
	fmt.Println(carro)

	carro2 := marcaCarros2(5)
	fmt.Println(carro2)

}
