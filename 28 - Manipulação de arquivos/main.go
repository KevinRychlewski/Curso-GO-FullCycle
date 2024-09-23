package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, err := os.Create("arquivo.txt")
	if err != nil {
		panic(err)
	}
	//Escrevendo dentro do arquivo
	tamanho, err := f.Write([]byte("Hello world"))
	//tamanho, err := f.WriteString("Hello world")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Arquivo criado com sucesso. Tamanho: %d bytes\n", tamanho)
	f.Close()

	// Leitura
	arquivo, err := os.ReadFile("arquivo.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(arquivo))

	// Leitura de pouco em pouco abrindo o arquivo
	arquivo2, err := os.Open("arquivo.txt")
	if err != nil {
		panic(err)
	}
	reader := bufio.NewReader(arquivo2)
	buffer := make([]byte, 3)
	for {
		n, err := reader.Read(buffer)
		if err != nil {
			break
		}
		fmt.Println(string(buffer[:n]))
	}

	// Removendo um arquivo
	err = os.Remove("arquivo.txt")
	if err != nil {
		panic(err)
	}
}