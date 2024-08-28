package main

import (
	"errors"
	"fmt"
)

func main() {

	// Como retornar mais de um valor, variavel e depois o outro valor
	valor, err := sum(50,10)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(valor)

	fmt.Println(sum(51, 2))
}

// Retornando mais de um valor
func sum(a, b int) (int,error) {
	if a + b >= 50 {
		return 0, errors.New("A soma é maior que 50")
	}
	return a + b, nil
}