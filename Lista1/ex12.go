// 12. Charrete
package main

import "fmt"

func main() {
	var h int

	fmt.Println("Digite a quantidade de horas de uso:")
	fmt.Scan(&h)

	valor := (h/3)*10 + (h%3)*5

	fmt.Println("O VALOR A PAGAR É = R$", valor)
}
