// Exercicío 18

package main

import "fmt"

func main() {
	var precoNormal float64
	var categoria string
	var diaSemana int

	fmt.Print("Digite o preço normal do DVD: R$")
	fmt.Scanln(&precoNormal)
	fmt.Print("Digite a categoria do DVD (comum ou lançamento): ")
	fmt.Scanln(&categoria)
	fmt.Print("Digite o dia da semana (1 para domingo, 2 para segunda, ..., 7 para sábado): ")
	fmt.Scanln(&diaSemana)

	var precoFinal float64
	switch diaSemana {
	case 2, 3, 5:
		precoFinal = precoNormal * 0.6 // Desconto de 40%
	default:
		precoFinal = precoNormal // Preço normal
	}
	fmt.Printf("O preço final do DVD é: R$ %.2f\n", precoFinal)

}
