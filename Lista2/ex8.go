// Exercicío 8

package main

import "fmt"

func main() {
	var n float64
	fmt.Println("Digite um número inteiro:")
	fmt.Scan(&n)

	if n > 20 && n < 90 {
		fmt.Printf("O número %v está compreendido entre 20 e 90", n)
	} else {
		fmt.Printf("O número %v não está compreendido entre 20 e 90", n)
	}
}
