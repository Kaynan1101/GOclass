package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

// Gerar - Vai retornar a aplicação de linha de comando pronta para ser executada
func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "Aplicação de linha de Comando"
	app.Usage = "Busca IP's e Nomes de Servidor da Internet"

	//seria o padrão do comando caso não passe nada no terminal a flag
	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "google.com.br",
		},
	}

	//comandos com suas respectivas funções
	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "Busca Ips de endereço na internet",
			Flags:  flags,
			Action: buscarIps,
		},
		{
			Name:   "servidores",
			Usage:  "Busca o nomes dos servidores na internet",
			Flags:  flags,
			Action: buscarServidores,
		},
	}

	return app

}

// função que vai buscar os Ips na internet
func buscarIps(c *cli.Context) {
	host := c.String("host")

	ips, erro := net.LookupIP(host)
	if erro != nil {
		log.Fatal(erro)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}

}

// função que vai buscar os servidores na internet
func buscarServidores(c *cli.Context) {
	host := c.String("host")
	servidores, erro := net.LookupNS(host)
	if erro != nil {
		log.Fatal(erro)
	}

	for _, servidor := range servidores {
		fmt.Println(servidor.Host)
	}
}
