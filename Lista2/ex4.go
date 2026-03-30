// Exercício 4
package main

import (
	"fmt"
	"math"
)

func main() {

	var numero float64

	fmt.Println("Digite um número:")
	fmt.Scan(&numero)

	if numero > 0 {
		fmt.Printf("Como o número é positivo, o resultado de sua raiz é %v", math.Sqrt(numero))
	} else if numero < 0 {
		fmt.Printf("Como o número é negativo, o resultado de sua potência é %v", math.Pow(numero, 2))
	}
}
