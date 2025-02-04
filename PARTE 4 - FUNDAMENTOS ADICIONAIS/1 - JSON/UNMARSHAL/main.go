package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type cachorro struct {
	Nome  string `json:"nome"`
	Raca  string `json:"raca"`
	Idade uint   `json:"idade"`
}

func main() {

	//JSON para STRUCT
	cachorroEmJSON := `{"nome":"Rex","raca":"Dálmata","idade":3}`

	//Para passar de JSON para struct cria uma variavel vazia
	var c cachorro

	//Unmarshal: precisa receber um slice de byte para converter
	//& para mostrar o endereço de memoria e alterar a variavel
	if erro := json.Unmarshal([]byte(cachorroEmJSON), &c); erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(c)

	//JSON para MAP
	cachorro2EmJSON := `{"nome":"Titan","raca":"Xiuaua"}`

	c2 := make(map[string]string)

	if erro := json.Unmarshal([]byte(cachorro2EmJSON), &c2); erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(c2)

}
