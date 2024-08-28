package main

import "fmt"

type Endereco struct {
	Logradouro string
	Numero int
	Cidade string
	Estado string
}

type cliente struct {
	Nome  string
	Idade int
	Ativo bool
	Endereco
}

func main() {
	kevin := cliente{
		Nome:  "Kevin",
		Idade: 30,
		Ativo: true,
	}
	kevin.Cidade = "São Paulo"

	fmt.Printf("Nome: %s, idade: %d, ativo: %t\n", kevin.Nome, kevin.Idade, kevin.Ativo)
}