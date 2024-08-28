package main

import "fmt"

func main() {

	salarios := map[string]int{"Kevin": 1000, "João": 2000, "Maria": 3000}
	fmt.Println(salarios["Kevin"])
	// Como deletar um item do map
	delete(salarios, "Kevin")
	// Como adicionar um item no map
	salarios["Lucas"] = 5000
	fmt.Println(salarios["Lucas"])

	// Jeitos de criar um map
	// salarios2 := make(map[string]int)
	// salarios1 := map[string]int{}
	// // Colocando valores dentro de um map
	// salarios1["Kevin"] = 1000

	// Percorrendo um map
	for nome, salario := range salarios {
		fmt.Printf("Nome: %s, Salario: %d\n",nome, salario)
	}

	// Como ignorar o valor de um map
	for _, salario := range salarios {
		fmt.Printf("Salario: %d\n", salario)
	}

}