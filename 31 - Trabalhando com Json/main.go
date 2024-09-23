package main

import (
	"encoding/json"
	"os"
)

type Conta struct {
	Numero int `json:"numero"`
	Saldo  int `json:"saldo"`
}

func main() {
	conta := Conta{Numero: 1, Saldo: 100}
	res, err := json.Marshal(conta)
	if err != nil {
		println(err)
	}
	println(string(res))

	// Criando um encoder
	err = json.NewEncoder(os.Stdout).Encode(conta)
	if err != nil {
		println(err)
	}

	// Criando um decoder
	jsonPuro := []byte(`{"numero": 2, "saldo": 200}`)
	var conta2 Conta
	err = json.Unmarshal(jsonPuro, &conta2)
	if err != nil {
		println(err)
	}
	println(conta2.Saldo)
}