// 17. Série Pares
package main

import "fmt"

func main() {
	var x, y int

	fmt.Println("Digite o primeiro número:")
	fmt.Scan(&x)

	fmt.Println("Digite a quantidade de termos:")
	fmt.Scan(&y)

	if x%2 != 0 {
		fmt.Println("O PRIMEIRO NUMERO NÃO É PAR")
	} else {
		fmt.Println("Serie gerada:")
		for i := 0; i < y; i++ {
			fmt.Print(x+i*2, " ")
		}
}
