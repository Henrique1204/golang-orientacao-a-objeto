package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func main() {
	primeiraConta := ContaCorrente{titular: "Paulo", numeroAgencia: 589, numeroConta: 123456, saldo: 323.45}

	fmt.Println(primeiraConta)

	var segundaConta *ContaCorrente
	segundaConta = new(ContaCorrente)

	fmt.Println(segundaConta)
}
