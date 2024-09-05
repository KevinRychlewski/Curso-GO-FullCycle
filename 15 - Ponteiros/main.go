package main

func main() {
	
	a := 10
	var ponteiro *int = &a
	// Mudando valor ponteiro
	*ponteiro = 20
	b := &a
	// Usar * na frente da variável para mostrar o valor do ponteiro e não o lugar da memória
	*b = 30
	println(*b)
}