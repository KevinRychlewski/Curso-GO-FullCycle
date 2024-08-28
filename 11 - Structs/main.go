package main

import "fmt"

type cliente struct {
	Nome  string
	Idade int
	Ativo bool
}

func main() {
	kevin := cliente{
		Nome:  "Kevin",
		Idade: 30,
		Ativo: true,
	}

	fmt.Printf("Nome: %s, idade: %d, ativo: %t\n", kevin.Nome, kevin.Idade, kevin.Ativo)
}