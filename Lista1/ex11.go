// 11. Divisível
package main

import "fmt"

func main() {
	var n int

	fmt.Println("Digite um número inteiro:")
	fmt.Scan(&n)

	if n%3 == 0 && n%5 == 0 {
		fmt.Println("Resultado: O NUMERO É DIVISIVEL POR 3 E 5")
	} else {
		fmt.Println("Resultado: O NUMERO NÃO É DIVISIVEL POR 3 E 5")
	}
}
