package main

import "fmt"

var numerador, denominador int
var somaH float64

func main() {
	numerador = 1
	denominador = 1
	for i := 1; i <= 50; i++ {
		somaH += float64(numerador) / float64(denominador)
		numerador += 2
		denominador += 1

	}
	fmt.Printf("A soma de H é de %.2f\n", somaH)
}
