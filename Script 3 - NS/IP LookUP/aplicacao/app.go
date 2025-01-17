package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
	// Gerar vai retornar a aplicação de linha de comando pronta para ser executada
)

func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "Aplicação de Linha de comando"
	app.Usage = "Busca IPs e Nomes de Servidor na Internet"

	app.Commands = []cli.Command{
		{
			// essa estrutura basicamente é uma struct
			Name:  "ip",
			Usage: "Busca IPs de endereços na internet",
			Flags: []cli.Flag{
				// aqui eu estou declarando uma nova struct que será atribuida a chave Flags, essa por sua vez vai receber os paramêtros via CLI como o Endereço de IP do servidor ou nome de domínio
				cli.StringFlag{ // essa struct irá receber os parametros e também irá definir os valores padrões caso nenhum parametro for passado.

					Name:  "host",
					Value: "www.udemy.com.br",
				},
			},
			Action: buscarIps, // usando a função action para executar a função buscarIps assim que os valores das chaves forem preenchidos.
		},
	} //slice de comandos que o programa poderá usar
	return app
}

func buscarIps(c *cli.Context) { //função buscar IP utilizando a biblioteca net para realizar as requisições e retornar as informações
	host := c.String("host") // passando o valor de host "Valor parametrizado ou padrão" para a variável host

	ips, erro := net.LookupIP(host) // usando a função LookupIP no host parametrizado ou no padrão para realizar uma busca
	// as linhas abaixo é para tratar o erro e printar os resultados.
	if erro != nil {
		log.Fatal(erro)
	}
	for _, ip := range ips {
		fmt.Println(ip)
	}

}
