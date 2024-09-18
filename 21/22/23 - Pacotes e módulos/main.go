package main

import (
	"fmt"
	"packages/matematica"
)

func main() {
	s := matematica.Soma(10,20)

	fmt.Printf("Resultado: %v", s)
}