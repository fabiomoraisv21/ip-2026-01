package main

import (
	"fmt"
	"strconv"
)

func pot(numero, expoente int) int {
	resultadoo := 1
	for i := 0; i < expoente; i++ {
		resultadoo *= numero
	}
	return resultadoo
}

var nBase8, somaExpo, p int

func main() {
	fmt.Println("Digite um número na base 8:")
	fmt.Scan(&nBase8)
	nbase8Str := strconv.Itoa(nBase8)
	lista := []int{}
	for i := len(nbase8Str) - 1; i >= 0; i-- {
		nBase8Int, _ := strconv.Atoi(string(nbase8Str[i]))
		lista = append(lista, nBase8Int)
	}
	fmt.Println(lista)
	for i := 0; i <= len(lista)-1; i++ {
		p = lista[i] * pot(8, i)
		fmt.Println(p)
		somaExpo += p

	}
	fmt.Printf("O numero na base 10 é %v\n", somaExpo)
}
