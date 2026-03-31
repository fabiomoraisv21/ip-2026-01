// 15. Quadrado dos Pares
package main

import "fmt"

func main() {
	var n int

	fmt.Println("Digite um número inteiro (5 < N < 2000):")
	fmt.Scan(&n)

	if n <= 5 || n >= 2000 {
		fmt.Println("Valor inválido!")
	} else {
		fmt.Println("Resultados:")
		for i := 2; i <= n; i += 2 {
			fmt.Println(i, "^2 =", i*i)
		}
	}
}
