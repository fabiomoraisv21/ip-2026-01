	// Exercicío 13

package main

import "fmt"

func main() {

	var n int
	fmt.Println("Digite um número inteiro positivo de 3 casas:")
	fmt.Scan(&n)

	if n >= 100 && n <= 999 {
		dezenas := (n / 10) % 10
		fmt.Printf("O algarismo da casa das dezenas é: %d", dezenas)
	} else {
		fmt.Println("O número informado não tem realmente 3 casas.")
	}

}
