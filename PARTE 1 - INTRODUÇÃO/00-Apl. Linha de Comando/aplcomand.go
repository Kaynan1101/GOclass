package main

import (
	"linha-de-comando/app"
	"log"
	"os"
)

// função main criada apenas para chamar o pacote app criado com os comando e funções passados
func main() {

	//chamando a função Gerar do pacoite app
	aplicacao := app.Gerar()

	//tratando erro de retorno
	// .RUN - maneira diferente de rodar um pacote externo
	if erro := aplicacao.Run(os.Args); erro != nil {
		log.Fatal(erro)
	}
}
