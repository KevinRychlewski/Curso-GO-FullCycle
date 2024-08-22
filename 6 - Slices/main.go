package main

import "fmt"

func main() {
	s := []int{10,20,30,40,50,60,70,80,90,100}
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s),s)
	// :0 significa que ele pega todos os valores da direita do slice e faz eles serem iguais a 0
	fmt.Printf("len=%d cap=%d %v\n", len(s[:0]), cap(s[:0]), s[:0])
	// :4 ele pega as 4 primeiras casas de um slice
	fmt.Printf("len=%d cap=%d %v\n", len(s[:4]), cap(s[:4]), s[:4])
	// 2: ele ignora os dois primeiros valores da slice e pega o restante
	fmt.Printf("len=%d caap=%d %v\n", len(s[2:]), cap(s[2:]), s[2:])

	// Aumentando o valor de um slice
	s = append(s, 110, 120, 130, 140, 150)
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)

}