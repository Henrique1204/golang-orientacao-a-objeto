package main

import (
	"fmt"

	"github.com/paulosouza/golang-orientacao-a-objeto/contas"
)

func main() {
	primeiraConta := contas.ContaCorrente{Titular: "Paulo", NumeroAgencia: 589, NumeroConta: 123456, Saldo: 323.45}

	fmt.Println(primeiraConta)

	segundaConta := new(contas.ContaCorrente)

	fmt.Println(segundaConta)

	fmt.Println(primeiraConta.Sacar(200))
	fmt.Println(primeiraConta.Sacar(200))
	fmt.Println(primeiraConta.Depositar(25.45))

	fmt.Println(primeiraConta)

	fmt.Println(primeiraConta.Transferir(100, segundaConta))

	fmt.Println(primeiraConta)
	fmt.Println(segundaConta)
}
