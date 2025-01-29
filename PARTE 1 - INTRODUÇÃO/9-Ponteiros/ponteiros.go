package main

import "fmt"

func main() {
	var var1 int = 15
	var ponteiro *int

	ponteiro = &var1
	fmt.Println(ponteiro)
	fmt.Println(*ponteiro)

}
