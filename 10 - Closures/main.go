package main

import (
	"fmt"
)

func main() {
	// Usando uma função anônima dentro da func main
	total := func() int{
		return sum(1,2,3,4,5,6,7,8,9,10) *2
	}()

	fmt.Println(total)

}

func sum(numeros ...int) int {
	total := 0
	for _, num := range numeros {
		total += num
	}
	return total
}