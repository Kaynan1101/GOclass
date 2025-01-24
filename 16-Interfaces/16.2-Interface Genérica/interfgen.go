package main

import "fmt"

func generica(interf interface{}) {
	fmt.Println(interf)
}
func main() {

	//recebe qualquer tipo de valor e não é passado metodo
	generica("STRING")
	generica(2)
	generica(false)

}
