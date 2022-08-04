package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func (c *ContaCorrente) Sacar(valorDoSaque float64) string {
	const MINIMO_SAQUE = 1

	podeSacar := valorDoSaque >= MINIMO_SAQUE && valorDoSaque <= c.saldo

	if podeSacar {
		c.saldo -= valorDoSaque

		return "Saque ocorreu com sucesso!"
	}

	return "Erro ao sacar!"
}

func main() {
	primeiraConta := ContaCorrente{titular: "Paulo", numeroAgencia: 589, numeroConta: 123456, saldo: 323.45}

	fmt.Println(primeiraConta)

	var segundaConta *ContaCorrente
	segundaConta = new(ContaCorrente)

	fmt.Println(segundaConta)

	fmt.Println(primeiraConta.Sacar(200))
	fmt.Println(primeiraConta.Sacar(200))
}
