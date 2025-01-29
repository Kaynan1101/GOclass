package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     uint8
	altura    float32
}

type estudante struct {
	pessoa     //'herança em go'
	curso      string
	falculdade string
}

func main() {

	fmt.Println("Herança")

	p1 := pessoa{"Kaynan", "Lima", 24, 1.78}
	fmt.Println(p1)

	e1 := estudante{p1, "ADS", "FATEC - PRESIDENTE PRUDENTE"}
	fmt.Println(e1)

}
