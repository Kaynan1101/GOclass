package main

import (
	"fmt"
)

func main() {
	var array = [2]int{2, 5}
	fmt.Println(array)

	array2 := [...]string{"Kaynan", "Lima", "de", "Oliveira"}
	fmt.Println(array2)

	slice := []int{1, 2, 3, 4, 5}
	fmt.Println(slice)

	slice = append(slice, 6)
	fmt.Println(slice)

	slice2 := array2[0:2]
	fmt.Println(slice2)

	array2[0] = "Fernanda"
	fmt.Println(slice2)

}
