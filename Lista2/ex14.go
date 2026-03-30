// Exercicío 14

package main

import "fmt"

func main() {

	var precoinicial float64
	var opcao string

	fmt.Println("Digite o preço inicial do carro:")
	fmt.Scan(&precoinicial)
	fmt.Println("Digite as opções desejadas (a, b, c, d):")
	fmt.Scan(&opcao)

	precofinal := precoinicial

	for _, letra := range opcao {
		switch letra {
		case 'a':
			precofinal += 1750
		case 'b':
			precofinal += 800
		case 'c':
			precofinal += 1200
		case 'd':
			precofinal += 2000
		default:
			fmt.Printf("Opção %v inválida!", string(letra))
		}
	}
	fmt.Printf("O preço final do carro é: R$ %.2f", precofinal)

}
