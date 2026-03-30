// Exercicio 5
package main

import "fmt"

func main() {
	var n int
	fmt.Println("Digite um número inteiro")
	fmt.Scan(&n)

	if n%5 == 0 {
		fmt.Printf("O número %v é divisível por 5", n)
	} else if n%5 != 0 {
		fmt.Printf("O número %v não é divisível por 5", n)
	}
}
