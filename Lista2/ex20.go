// Exercicío 20

package main

import "fmt"

func main() {
	var preco float64
	var opcao int

	fmt.Print("Digite o preço do produto: ")
	fmt.Scanln(&preco)	
	fmt.Println("Escolha a condição de pagamento:")
	fmt.Println("1 - À vista, dinheiro ou cheque (10% de desconto)")
	fmt.Println("2 - À vista, cartão de crédito (5% de desconto)")
	fmt.Println("3 - Em 2 vezes (preço normal de etiqueta sem juros)")
	fmt.Println("4 - Em 3 vezes (preço normal de etiqueta + 10% de juros)")
	fmt.Scanln(&opcao)

	switch opcao {
	case 1:
		precoFinal := preco * 0.90
		fmt.Printf("Preço final: %.2f\n", precoFinal)
	case 2:
		precoFinal := preco * 0.95
		fmt.Printf("Preço final: %.2f\n", precoFinal)
	case 3:
		precoFinal := preco
		fmt.Printf("Preço final: %.2f\n", precoFinal)
	case 4:
		precoFinal := preco * 1.10
		fmt.Printf("Preço final: %.2f\n", precoFinal)	
	default:
		fmt.Println("Opção inválida!")
	}	
}