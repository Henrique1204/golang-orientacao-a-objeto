package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func (c *ContaCorrente) Sacar(valorDoSaque float64) (string, float64) {
	const SAQUE_MINIMO = 0

	podeSacar := valorDoSaque > SAQUE_MINIMO && valorDoSaque <= c.saldo

	if podeSacar {
		c.saldo -= valorDoSaque

		return "Saque ocorreu com sucesso!", c.saldo
	}

	return "Erro ao sacar!", c.saldo
}

func (c *ContaCorrente) Depositar(valorDoDeposito float64) (string, float64) {
	const MINIMO_DEPOSITO = 0

	if valorDoDeposito > MINIMO_DEPOSITO {
		c.saldo += valorDoDeposito

		return "Deposito ocorreu com sucesso!", c.saldo
	}

	return "Erro ao depositar!", c.saldo
}

func main() {
	primeiraConta := ContaCorrente{titular: "Paulo", numeroAgencia: 589, numeroConta: 123456, saldo: 323.45}

	fmt.Println(primeiraConta)

	var segundaConta *ContaCorrente
	segundaConta = new(ContaCorrente)

	fmt.Println(segundaConta)

	fmt.Println(primeiraConta.Sacar(200))
	fmt.Println(primeiraConta.Sacar(200))
	fmt.Println(primeiraConta.Depositar(25.45))

	fmt.Println(primeiraConta)
}
