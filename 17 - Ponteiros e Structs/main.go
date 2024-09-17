package main

import "fmt"

type Conta struct {
	saldo float64
}

func newConta() *Conta {
	return &Conta{saldo: 0}
}

func (c *Conta) simular(valor float64) float64{
	c.saldo += valor
	return c.saldo
}

func main() {
	conta := Conta{
		saldo: 100,
	}
	conta.simular(200)

	fmt.Println(conta.saldo)
}


