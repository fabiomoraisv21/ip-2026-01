// 19. Somatório
package main

import "fmt"

func main() {
	var n int

	fmt.Println("Digite um número inteiro maior que 1:")
	fmt.Scan(&n)

	if n <= 1 {
		fmt.Println("Numero invalido!")
	} else {
		var s float64

		for i := 1; i <= n; i++ {
			s += 1.0 / float64(i)
		}

		fmt.Println("RESULTADO =", s)
	}
}
