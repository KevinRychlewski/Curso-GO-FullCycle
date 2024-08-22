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

	// %T mostra o tipo da variável
	fmt.Printf("o tipo de E é %T", e)
	// %v mostra o valor da variável
	fmt.Println("o tipo de E é %v", e)
}