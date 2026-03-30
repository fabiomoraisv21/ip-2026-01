// Exercício 2

package main

import "fmt"

func main() {

	var numero int
	fmt.Println("Digite um número inteiro:")
	fmt.Scan(&numero)

	if numero > 0 {
		fmt.Printf("O número %v é positivo", numero)
	} else if numero < 0 {
		fmt.Printf("O número %v é negativo", numero)
	} else {
		fmt.Printf("O número %v é nulo", numero)
	}
}
