package main

import (
	"fmt"
)

// 1. Média Semestral
func mediaSemestral() {
	var m1, m2, m3, mf float64

	fmt.Print("Digite a primeira nota: ")
	fmt.Scan(&m1)
	fmt.Print("Digite a segunda nota: ")
	fmt.Scan(&m2)
	fmt.Print("Digite a terceira nota: ")
	fmt.Scan(&m3)

	mf = (m1 + m2 + m3) / 3

	if mf >= 6 {
		fmt.Println("Aprovado com média:", mf)
	} else {
		fmt.Println("Reprovado com média:", mf)
	}
}
