package main

import (
	"banco/contas"
	"fmt"
)

func main() {
	contaDenis := contas.ContaPoupanca{}
	contaDenis.Depositar(100)
	contaDenis.Sacar(25)

	fmt.Println(contaDenis.ObterSaldo())

	PagarBoleto(&contaDenis, 30)
	fmt.Println(contaDenis.ObterSaldo())
}

func PagarBoleto(conta VerificarConta, valorBoleto float64) {
	conta.Sacar(valorBoleto)
}

type VerificarConta interface {
	Sacar(valor float64) string
}
