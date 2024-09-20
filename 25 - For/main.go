package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		println(i)
	}

	numeros := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, num := range numeros {
		println(num)
	}

	i := 0
	for i < 10 {
		println(i)
		i++
	}
	// Loop infinito
	for {
		fmt.Println("Hello, World")
	}
}