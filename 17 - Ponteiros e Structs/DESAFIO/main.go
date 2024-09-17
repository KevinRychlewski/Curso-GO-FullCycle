package main

import "fmt"

type Pessoa struct {
	nome string
	idade int
	endereco string
}

func (p *Pessoa) imprimirPessoa() {
	fmt.Printf("O nome da pessoa é %s, a idade é %d e o endereço é %s\n", p.nome, p.idade, p.endereco)
}

func (p *Pessoa) alterarNome(novoNome string) {
	p.nome = novoNome
}

func (p *Pessoa) alterarIdade(novaIdade int) {
	p.idade = novaIdade
}

func (p *Pessoa) alterarEndereco(novoEndereco string) {
	p.endereco = novoEndereco
}

func main() {
	pessoa := Pessoa {
		nome: "Kevin",
		idade: 18,
		endereco: "Avenida dos idiotas",
	}

	pessoa.imprimirPessoa()
	pessoa.alterarNome("Pedro")
	pessoa.alterarIdade(20)
	pessoa.alterarEndereco("Rua dos idiotas")
	pessoa.imprimirPessoa()
}