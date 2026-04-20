package main

import (
	"fmt"
	"strconv"
)

var cpf, cpf2 string
var somaCPF, primeiroDigito, somacpf2, segundoDigito int

func main() {
	fmt.Println("Digite seu CPF:")
	fmt.Scan(&cpf)
	multiplicador := 10
	for i := 0; i < 9; i++ {
		digito, _ := strconv.Atoi(string(cpf[i]))
		digito *= multiplicador
		somaCPF += digito
		multiplicador--
	}
	restoCpf := somaCPF % 11
	if restoCpf < 2 {
		primeiroDigito = 0
	} else {
		primeiroDigito = 11 - restoCpf
	}
	cpf2 = string(cpf[0:9]) + strconv.Itoa(primeiroDigito)
	multiplicador = 11
	for i := 0; i <= 9; i++ {
		digito, _ := strconv.Atoi(string(cpf2[i]))
		digito *= multiplicador
		somacpf2 += digito
		multiplicador--
	}
	restoCpf2 := somacpf2 % 11
	if restoCpf2 < 2 {
		segundoDigito = 0
	} else {
		segundoDigito = 11 - restoCpf2
	}
	primeiroStr := strconv.Itoa(primeiroDigito)
	segundoStr := strconv.Itoa(segundoDigito)
	if primeiroStr == string(cpf[9]) && segundoStr == string(cpf[10]) {
		fmt.Println("Esse cpf é valido")
	} else {
		fmt.Println("Esse cpf não é valido")
	}

}
