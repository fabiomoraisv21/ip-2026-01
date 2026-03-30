// Exercicío 7
package main

import "fmt"

func main() {
	var A, B, C int
	var MENOR, INTER, MAIOR int

	fmt.Println("Digite os números A, B e C (espaçados):")
	fmt.Scan(&A, &B, &C)

	if A > B && A > C {
		MAIOR = A
	} else if B > A && B > C {
		MAIOR = B
	} else {
		MAIOR = C
	}

	if A < B && A < C {
		MENOR = A
	} else if B < A && B < C {
		MENOR = B
	} else {
		MENOR = C
	}

	INTER = (A + B + C) - MAIOR - MENOR

	fmt.Printf("A ordem crescente é: %d, %d, %d\n", MENOR, INTER, MAIOR)
}
