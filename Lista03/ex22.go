package main

import "fmt"

var numerador, denominador, soma int

func main() {
	numerador = 38
	denominador = 1
	for i := 1; i <= 37; i++ {
		soma += (numerador * (numerador - 1)) / denominador
		numerador -= 1
		denominador += 1
	}
	fmt.Println("O resultado é", soma)
}
