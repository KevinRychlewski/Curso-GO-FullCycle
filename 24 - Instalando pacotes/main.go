package main

import (
	"fmt"
	"packages/matematica"
	"github.com/google/uuid"
)

func main() {
	s := matematica.Soma(10,20)
	carro := matematica.Carro{Marca: "Ford"}
	fmt.Printf("Resultado: %v", s)
	fmt.Println(matematica.A)
	fmt.Println(carro)
	fmt.Println(carro.Andar())
	fmt.Println(uuid.New())
}