// Exercicío 3 
package main

import "fmt"

func main() {

	var n1, n2 float64
	fmt.Println("Digite o primeiro número:")
	fmt.Scan(&n1)
	fmt.Println("Digite o segundo número:")
	fmt.Scan(&n2)

	if n1+n2 > 20 {
		fmt.Print("O resultado é ", n1+n2+8)
	} else if n1+n2 <= 20 {
		fmt.Print("O resultado é ", n1+n2-5)
	}
}
