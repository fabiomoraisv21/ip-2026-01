// Exercicío 15

package main

import "fmt"

func main() {
	var dia, mes, ano int
	fmt.Println("Digite a data (dia, mês e ano, separados por espaço):")
	fmt.Scan(&dia, &mes, &ano)

	switch mes {
	case 1:
		fmt.Printf("%d de Janeiro de %d", dia, ano)
	case 2:
		fmt.Printf("%d de Fevereiro de %d", dia, ano)
	case 3:
		fmt.Printf("%d de Março de %d", dia, ano)
	case 4:
		fmt.Printf("%d de Abril de %d", dia, ano)
	case 5:
		fmt.Printf("%d de Maio de %d", dia, ano)
	case 6:
		fmt.Printf("%d de Junho de %d", dia, ano)
	case 7:
		fmt.Printf("%d de Julho de %d", dia, ano)
	case 8:
		fmt.Printf("%d de Agosto de %d", dia, ano)
	case 9:
		fmt.Printf("%d de Setembro de %d", dia, ano)
	case 10:
		fmt.Printf("%d de Outubro de %d", dia, ano)
	case 11:
		fmt.Printf("%d de Novembro de %d", dia, ano)
	case 12:
		fmt.Printf("%d de Dezembro de %d", dia, ano)
	default:
		fmt.Println("Mês inválido!")
	}
}
