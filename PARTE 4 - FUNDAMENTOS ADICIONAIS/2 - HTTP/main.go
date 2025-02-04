package main

import (
	"log"
	"net/http"
)

// PADRÃO DA ROTA
func home(w http.ResponseWriter, r *http.Request) {
	//NÃO UTILIZA FMT.PRINTLN NO HTTP
	w.Write([]byte("Olá Mundo!"))
}

func usuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Cadastrando Usuários"))
}

func main() {
	//HTTP = REQUEST e RESPONSE
	//ROTAS:
	//-URI: IDENTIFICADOR DO RECURSO
	//-MÉTODO: GET, POST, PUT E DELETE

	//ROTA
	http.HandleFunc("/home", home)
	http.HandleFunc("/usuarios", usuario)

	//ABRINDO UM SERVIDOR
	log.Fatal(http.ListenAndServe(":5000", nil))

}
