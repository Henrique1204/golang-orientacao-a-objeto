package contas

import (
	"github.com/paulosouza/golang-orientacao-a-objeto/clientes"
)

type ContaCorrente struct {
	Titular       clientes.Titular
	NumeroAgencia int
	NumeroConta   int
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
		return status
	}

	status, _ = contaDeDestino.Depositar(valorDaTransferencia)

	if !status {
		c.Depositar(valorDaTransferencia)

		return status
	}

	return status
}

func (c *ContaCorrente) ConsultarSaldo() float64 {
	return c.saldo
}
