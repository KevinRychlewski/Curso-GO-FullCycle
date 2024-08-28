package main

import (
	"fmt"
)

func main() {

	fmt.Println(sum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))

}

// Função variádica
func sum(numeros ...int) int {
	total := 0
	for _, num := range numeros {
		total += num
	}
	return total
}