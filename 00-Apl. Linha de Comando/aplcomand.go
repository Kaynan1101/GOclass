package main

import (
	"linha-de-comando/app"
	"log"
	"os"
)

func main() {

	aplicacao := app.Gerar()

	//tratando erro de retorno
	if erro := aplicacao.Run(os.Args); erro != nil {
		log.Fatal(erro)
	}
}
