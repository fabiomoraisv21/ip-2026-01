// Exercicío 9

package main

import "fmt"

func main() {
	var valorCompra float64
	fmt.Println("Digite o valor da compra:")
	fmt.Scan(&valorCompra)
	if valorCompra < 0 {
		fmt.Println("Valor de compra inválido!")
	} else if valorCompra < 10 {
		fmt.Printf("O valor da venda é R$ %v", valorCompra*1.7)
	} else if valorCompra >= 10 && valorCompra < 30 {
		fmt.Printf("O valor da venda é R$ %v", valorCompra*1.5)
	} else if valorCompra >= 30 && valorCompra < 50 {
		fmt.Printf("O valor da venda é R$ %v", valorCompra*1.4)
	} else if valorCompra >= 50 {
		fmt.Printf("O valor da venda é R$ %v", valorCompra*1.3)
	}
}
