package main

import (
	"database/sql"
	"fmt"
	"log"

	//Importando pacotes implicíto
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	//string de conexão padrão para se conectar ao banco MYSQL, se fosse outro banco
	//como postgre usaria outra string de conexão
	stringConexao := "golang:golang@/devbook?charset=utf8&parseTime=True&loc=Local"
	db, erro := sql.Open("mysql", stringConexao)
	if erro != nil {
		log.Fatal(erro)
	}

	defer db.Close() //Fecha a conexão com banco

	//Testa a conexão com o banco
	if erro = db.Ping(); erro != nil {
		log.Fatal(erro)
	}

	fmt.Println("Conexão Aberta")

	linhas, erro := db.Query("select * from usuarios")
	if erro != nil {
		log.Fatal(erro)
	}

	defer linhas.Close()

	fmt.Println(linhas)

}
