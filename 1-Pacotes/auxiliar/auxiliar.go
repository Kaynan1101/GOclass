package auxiliar

import "fmt"

//Escrever aparece um print
func Escrevendo() {
	fmt.Println("Escrevendo do arquivo main")
	//não é preciso importar o modulo porque eles estão dentro do mesmo pacote
	escrevendo() //como está dentro da mesma função ele irá aparecer na execução da main.go, direto na main já não consigo

}
