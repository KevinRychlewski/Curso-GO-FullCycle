package main

import "fmt"

const a = "Hello, World!"

type ID int

var(
b bool = true
c int = 10
d string = "Kevin"
e float64 = 1.2
f ID = 1
)

func main() {
	// Declarando um array
	var meuArray [3]int
	meuArray[0] = 10
	meuArray[1] = 20
	meuArray[2] = 30
	// Imprimindo o tamanho do array
	fmt.Println(len(meuArray))
	// Acessando o ultimo item do array
	fmt.Println(meuArray[len(meuArray)-1])
	// Percorrendo pelo array
	for indice, valor := range meuArray {
		fmt.Printf("o valor no indice é %d e o valor é %d\n", indice, valor)
	}
}