package main

import "fmt"

type Endereco struct {
	Logradouro string
	Numero int
	Cidade string
	Estado string
}

type Pessoa interface {
	Desativar()
}

type Empresa struct{
	Nome string
}

func (e Empresa) Desativar() {

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

func Desativacao(pessoa Pessoa) {
	pessoa.Desativar()
}

func main() {
	kevin := cliente{
		Nome:  "Kevin",
		Idade: 30,
		Ativo: true,
	}
	minhaEmpresa := Empresa{
	}
	Desativacao(minhaEmpresa)
}