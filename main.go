package main

import (
	"fmt"

	"github.com/paulosouza/golang-orientacao-a-objeto/clientes"
	"github.com/paulosouza/golang-orientacao-a-objeto/contas"
)

type verificarConta interface {
	Sacar(valor float64) (bool, float64)
}

func PagarBoleto(conta verificarConta, valorDoBoleto float64) bool {
	status, _ := conta.Sacar(valorDoBoleto)

	return status
}

func main() {
	paulo := clientes.Titular{Nome: "Paulo Souza", CPF: "123.456.789.00", Profissao: "Desenvolvedor"}

	contaDoPaulo := contas.ContaCorrente{Titular: paulo, NumeroAgencia: 001, NumeroConta: 123}
	contaDoPaulo.Depositar(200)

	fmt.Println(contaDoPaulo.ConsultarSaldo())

	henrique := clientes.Titular{Nome: "Henrique Souza", CPF: "123.456.789.00", Profissao: "Desenvolvedor"}

	contaDoHenrique := contas.ContaPoupanca{Titular: henrique, NumeroAgencia: 001, NumeroConta: 123}
	contaDoHenrique.Depositar(200)

	fmt.Println(contaDoHenrique.ConsultarSaldo())

	PagarBoleto(&contaDoHenrique, 100)
	PagarBoleto(&contaDoPaulo, 100)
}
