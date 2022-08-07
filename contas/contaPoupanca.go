package contas

import (
	"github.com/paulosouza/golang-orientacao-a-objeto/clientes"
)

type ContaPoupanca struct {
	Titular                              clientes.Titular
	NumeroAgencia, NumeroConta, Operacao int
	saldo                                float64
}

func (c *ContaPoupanca) Sacar(valorDoSaque float64) (bool, float64) {
	const SAQUE_MINIMO = 0

	podeSacar := valorDoSaque > SAQUE_MINIMO && valorDoSaque <= c.saldo

	if podeSacar {
		c.saldo -= valorDoSaque
	}

	return podeSacar, c.saldo
}

func (c *ContaPoupanca) Depositar(valorDoDeposito float64) (bool, float64) {
	const MINIMO_DEPOSITO = 0

	podeDepositar := valorDoDeposito > MINIMO_DEPOSITO

	if podeDepositar {
		c.saldo += valorDoDeposito
	}

	return podeDepositar, c.saldo
}

func (c *ContaPoupanca) Transferir(valorDaTransferencia float64, contaDeDestino *ContaPoupanca) bool {
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

func (c *ContaPoupanca) ConsultarSaldo() float64 {
	return c.saldo
}
