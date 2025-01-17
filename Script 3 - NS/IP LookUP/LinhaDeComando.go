package main

import (
	"fmt"
	app "linha-de-comando/aplicacao"
	"log"
	"os"
)

func main() {
	fmt.Println("Ola Mundo")
	aplicacao := app.Gerar()
	erro := aplicacao.Run(os.Args) // Esse argumento é util para que os comandos que serem passados pelo prompt de comando sejam reconhecidos durante a execução.
	if erro != nil {               // tratando o erro do .run
		log.Fatal(erro)
	}
}
