// Exercicío 11

package main

import "fmt"

func main() {

	var x float64
	fmt.Println("Digite o valor de x:")
	fmt.Scan(&x)

	if x != 2 {
		fmt.Printf("O resultado de f(x) é %v", 8/(2-x))
	} else {
		fmt.Println("Valor inválido para x!")
	}
}
