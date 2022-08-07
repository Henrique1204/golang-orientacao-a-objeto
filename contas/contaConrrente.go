package contas

import "fmt"

type ContaCorrente struct {
	Titular       string
	NumeroAgencia int
	NumeroConta   int
	Saldo         float64
}

func (c *ContaCorrente) Sacar(valorDoSaque float64) (bool, float64) {
	const SAQUE_MINIMO = 0

	podeSacar := valorDoSaque > SAQUE_MINIMO && valorDoSaque <= c.Saldo

	if podeSacar {
		c.Saldo -= valorDoSaque
	}

	return podeSacar, c.Saldo
}

func (c *ContaCorrente) Depositar(valorDoDeposito float64) (bool, float64) {
	const MINIMO_DEPOSITO = 0

	podeDepositar := valorDoDeposito > MINIMO_DEPOSITO

	if podeDepositar {
		c.Saldo += valorDoDeposito
	}

	return podeDepositar, c.Saldo
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
