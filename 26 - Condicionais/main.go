package main

func main() {
	a := 1
	b := 2
	c := 3

	switch a {
	case 1:
		println("a")
	case 2:
		println("b")
	case 3:
		println("c")
	// Caso não seja nenhuma das opções ele executa o default
	default:
		println("d")
	}


	// Condicional E (&&)
	if a > b && a > c {
		println("a é maior")
	}

	// Condicional OU (||)
	if a > b || a > c {
		println("a é maior")
	}
 
	if a > b {
		println(a)
	} else {
		println(b)
	}
}
