package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func (c *ContaCorrente) Sacar(valorDoSaque float64) (bool, float64) {
	const SAQUE_MINIMO = 0

	podeSacar := valorDoSaque > SAQUE_MINIMO && valorDoSaque <= c.saldo

	if podeSacar {
		c.saldo -= valorDoSaque
	}

	return podeSacar, c.saldo
}

func (c *ContaCorrente) Depositar(valorDoDeposito float64) (bool, float64) {
	const MINIMO_DEPOSITO = 0

	podeDepositar := valorDoDeposito > MINIMO_DEPOSITO

	if podeDepositar {
		c.saldo += valorDoDeposito
	}

	return podeDepositar, c.saldo
}

func (c *ContaCorrente) Transferir(valorDaTransferencia float64, contaDeDestino *ContaCorrente) bool {
	status, _ := c.Sacar(valorDaTransferencia)

	if !status {
		fmt.Println("Você não tem saldo suficiente para essa transferência!")

		return status
	}

	status, _ = contaDeDestino.Depositar(valorDaTransferencia)

	if !status {
		fmt.Println("Ocorreu um erro ao depositar na conta de destino!")

		c.Depositar(valorDaTransferencia)

		return status
	}

	return status

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

	fmt.Println(primeiraConta.Transferir(100, segundaConta))

	fmt.Println(primeiraConta)
	fmt.Println(segundaConta)
}
