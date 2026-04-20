package main

import "fmt"

var termos int

func main() {
	primeiroTermo := 0
	segundoTermo := 1
	listaFibonacci := []int{}
	fmt.Println("Quantos termos da sequencia de Fibonacci?")
	fmt.Scan(&termos)
	for i := 0; i < termos; i++ {
		listaFibonacci = append(listaFibonacci, primeiroTermo)
		primeiroTermo, segundoTermo = segundoTermo+primeiroTermo, primeiroTermo
	}
	fmt.Println(listaFibonacci)
}
