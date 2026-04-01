package main

import "fmt"

func main() {

	var notas [3]float64
	var soma float64

	for i := 0; i < 3; i++ {

		fmt.Printf("Digite a nota %d:", i+1)
		fmt.Scan(&notas[i])

		soma += notas[i]
	}
	mediag := soma / 3

	fmt.Printf("A média do aluno é: %.2f", mediag)
}
