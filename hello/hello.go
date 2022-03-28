package main

import (
	"fmt"
)

func main() {
	nome := "Bruna"
	versao := 1.1

	fmt.Println("Olá,", nome)
	fmt.Println("Este programa está na versão:", versao)

	fmt.Println("1 - Iniciar monitoramento")
	fmt.Println("2 - Exibir logs")
	fmt.Println("0 - Sair")

	var comando int
	fmt.Scan(&comando)
	fmt.Println("Comando informado: ", comando)
	fmt.Println("Endereço da variável comando:", &comando)
}
