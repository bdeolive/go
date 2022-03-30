package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func main() {
	contaBruna := ContaCorrente{"Bruna", 589, 123456, 125.5}

	fmt.Println(contaBruna.titular)
}
