package main

import (
	"fmt"

	"github.com/paulosouza/golang-orientacao-a-objeto/clientes"
	"github.com/paulosouza/golang-orientacao-a-objeto/contas"
)

func main() {
	paulo := clientes.Titular{Nome: "Paulo Souza", CPF: "123.456.789.00", Profissao: "Desenvolvedor"}

	contaDoPaulo := contas.ContaCorrente{Titular: paulo, NumeroAgencia: 001, NumeroConta: 123, Saldo: 500}

	fmt.Println(contaDoPaulo)
}
