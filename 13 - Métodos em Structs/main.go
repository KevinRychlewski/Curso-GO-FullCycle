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

// Como criar um método em uma struct
func (c *cliente) Desativar() {
	c.Ativo = false
	fmt.Printf("O cliente %s foi desativado\n", c.Nome)
}

func main() {
	kevin := cliente{
		Nome:  "Kevin",
		Idade: 30,
		Ativo: true,
	}
	kevin.Ativo = false
	kevin.Desativar()
}