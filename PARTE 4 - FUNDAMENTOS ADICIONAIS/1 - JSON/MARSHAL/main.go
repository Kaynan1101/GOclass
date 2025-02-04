package main

import (
	"bytes"
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

	c := cachorro{"Rex", "Dálmata", 3}

	//Convertando uma struct em JSON.
	//Em GO tudo que for em JSON, sempre deixar na letra maiuscula por padrão
	//json.Marshal realiza a conversao
	//json.Umarshal converte o JSON em struct
	cachorroEmJSON, erro := json.Marshal(c)
	if erro != nil {
		log.Fatal(erro)
	}

	//Devolve um slice de bytes
	fmt.Println(cachorroEmJSON)
	//Converte o slice de bytes em json legivel
	fmt.Println(bytes.NewBuffer(cachorroEmJSON))

	//Convertendo um Map
	c2 := map[string]string{
		"nome": "Titan",
		"raca": "Xiuaua",
	}

	cachorro2EmJSON, erro := json.Marshal(c2)
	if erro != nil {
		log.Fatal(erro)
	}

	//Devolve um slice de bytes
	fmt.Println(cachorro2EmJSON)
	//Converte o slice de bytes em json legivel
	fmt.Println(bytes.NewBuffer(cachorro2EmJSON))

}
