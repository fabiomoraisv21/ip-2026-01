// Exercicío 16

package main

import (
	"fmt"
)

func main() {
	var A, B, C float64
	fmt.Println("Digite os coeficientes A, B e C (espaçados):")
	fmt.Scan(&A, &B, &C)

	if A == 0 {
		fmt.Println("Não é uma equação do segundo grau!")
		return
	}

	delta := B*B - 4*A*C

	if delta < 0 {
		fmt.Println("RAÍZES IMAGINÁRIAS")
	} else if delta == 0 {
		fmt.Println("RAIZ ÚNICA")
	} else {
		fmt.Println("RAÍZES DISTINTAS")
	}

}
