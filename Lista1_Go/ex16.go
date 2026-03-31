// 16. Reajuste
package main

import "fmt"

func main() {
	var salario float64

	fmt.Println("Digite o salário atual:")
	fmt.Scan(&salario)

	var novo float64

	if salario <= 300 {
		novo = salario * 1.5
	} else {
		novo = salario * 1.3
	}

	fmt.Println("SALÁRIO COM REAJUSTE = R$", novo)
}
